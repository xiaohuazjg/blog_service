package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/internal/services"
	"github.com/xiaohuazjg/blog_service/pkg/app"
	"github.com/xiaohuazjg/blog_service/pkg/convert"
	"github.com/xiaohuazjg/blog_service/pkg/errcode"
	"github.com/xiaohuazjg/blog_service/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	reponse := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		reponse.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	filetype := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || filetype <= 0 {
		reponse.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := services.New(c.Request.Context())
	fileinfo, err := svc.UploadFile(upload.FIleType(filetype), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err:%v", err)
		reponse.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	reponse.ToResponse(gin.H{
		"file_access_url": fileinfo.AccessUrl,
	})

}
