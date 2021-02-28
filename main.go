package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/internal/model"
	"github.com/xiaohuazjg/blog_service/internal/routers"
	"github.com/xiaohuazjg/blog_service/pkg/logger"
	"github.com/xiaohuazjg/blog_service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err%v")
	}
	err = setupLogger()
	if err != nil {
		log.Fatal("init.setupLogger err:", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("server", &global.ServerSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return nil
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

//@title 博客系统
//@version 0.1
//@description Go 语言编程之旅
//@teams of service github.com/xiaohuazjg/blog_service
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	filename := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filename,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
