package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/pkg/app"
	"github.com/xiaohuazjg/blog_service/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err:%v", err)
				err := defailtMailer.SendMail(
					global.EmailSetting.TO,
					fmt.Printf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Printf("错误信息：%v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err:%v", err)

				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Next()
			}
		}()
		c.Next()
	}
}
