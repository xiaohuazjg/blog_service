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
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"gopkg.in/natefinch/lumberjack.v2"
)


var (
	port string
	runMode string
	config string
	isVersion bool

)
func init() {

	err:=setupFlag()
	if err!=nil {
		log.Fatalf("init.setupFlag err:v%",err)
	}

	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v",err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatal("init.setupLogger err:%v", err)
	}
	err=setupDBEngine(){
       if err!=nil{
		log.Fatal("init.setupDBEngine err:%v", err)
	   }
	}
	err=setupValidator(){
		if err!=nil{
		 log.Fatal("init.setupValidator err:%v", err)
		}
	 }
	 err=setupTracer(){
		if err!=nil{
		 log.Fatal("init.setupValidator err:%v", err)
		}
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
   go func(){
	   if err:=s.ListenAndServe();err!=nil &&err!=http.ErrServerClosed {
		   log.Fatalf("s.ListenAndServer() err:%v")
	   }
   }()
   quit:=make(char os.Signal)
   singnal.Notify(quit,syscall.SIGINT,syscall.SYSTEM)
   <-quit
   log.Println("Shuting downing server...")
   
   ctx,cancel :=context.WithTimeout(context,Background(),5*time.Second)
	defer cancel()
	if err:=s.Shutdown(ctx);err!=nil {
		log.Fatalf("Server forced shutdwon:",err)

	}
     log.Println("Server Exiting")
}

func setupFlag() error {
	flag.StringVar(&port,"port","","启动端口")
	flag.StringVar(&runMode,"mode","","启动模式")
	flag.StringVar(&config,"config","configs/","指定使用的配置文件路径")
	flag.BoolVar(&isVersion,"version",false,"编译信息")
    flag.Parse()
    return nil 

}

func setupSetting() error {
   s,err:=setting.NewSetting(stings.Split(config,",")...)
   if err!=nil {
	   return err

   }
   err=s.ReadSection("Server",&global.ServerSetting)
   
	   if err!=nil {
		   return err
	   }
   
   err=s.ReadSection("App",&global.AppSetting)
	if err!=nil {
		return err
	}
	err=s.ReadSection("Database",&global.DatabaseSetting)
	if err!=nil {
		return err
	}
	err=s.ReadSection("JWT",&global.JWTSetting)
	if err!=nil {
		return err
	}
	err=s.ReadSection("Email",&global.EmailSetting)
	if err!=nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout*=time.Second
	global.JWTSetting.Expire*=time.Second
	global.ServerSetting.ReadTimeout*=time.Second
	global.ServerSetting.WriteTimeout*=time.Second
	
	if port!="" {
		global.ServerSetting.HttpPort=port

	}
	if runMode!="" {
		global.ServerSetting.RunMode=runMode
	}
	return nil 

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


func setupTracer() error {
    jaegerTracer,_,err:=tracer.NewJadgerTrader("blog_service","127.0.0.1:6831")
	if err!=nil {
		return err
	}
	global.Tracer =jaegerTracer
	return nil
	
}