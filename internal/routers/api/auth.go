package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/internal/service"
	"github.com/xiaohuazjg/blog_service/pkg/app"
	"github.com/xiaohuazjg/blog_service/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid err:v%", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return

	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err:%v", err)
		response.ToErrorResponse(errcode.UnathorizedAuthNotExist)
		return

	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokeGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
