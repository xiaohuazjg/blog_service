package global

import (
	"github.com/xiaohuazjg/blog_service/pkg/logger"
	"github.com/xiaohuazjg/blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
