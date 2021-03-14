package routers

import (
	"/blog_service/routers/api/v1/article"
	"/blog_service/routers/api/v1/tag"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/internal/routers/api"
)
var methodLimiters=limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:"/auth",
		FillInterval:time.Second,
		Capacity:10,
		Quantum:10,
	},
)


func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode=="debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiters)
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	article:=v1.NewArticle()
	tag:=v1.NewTag()
	upload:=api.NewUpload()
	r.Get("/debug/vars",api.Expvar)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", api.GetAuth)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles", article.List)

	}
	return r

}
