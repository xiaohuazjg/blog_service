package model

import (
	"fmt"
	"time"

	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/pkg/setting"
	"gorm.io/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreateBy   string `json:"create_by"`
	CreateOn   uint32 `json:"create_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {

	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.logMode(true)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Create().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil

}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreateOn"); ok {
			if createTimeField.IsBlank() {
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank() {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); ok {
		_ = scope.SetColumn("ModifyOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_column"); ok {
			extraOption = fmt.Sprintf(str)

		}
		deletedOnField, hasDeletedField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelfield := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedField && hasIsDelfield {
			nows := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quoted(deletedOnField.DBName),
				scope.AddToVars(nows),
				scope.Quoted(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombineConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM  %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombineConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
