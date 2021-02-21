package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/internal/routers"
	"github.com/xiaohuazjg/blog_service/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err%v")
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("server", global.ServerSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("App", global.AppSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("Database", global.DatabaseSetting)
	if err != nil {
		return nil
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxReadBytes: 1 << 20,
	}
	s.ListenAndServe()

}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabseSetting)
	if err != nil {
		return err
	}
	return nil
}
