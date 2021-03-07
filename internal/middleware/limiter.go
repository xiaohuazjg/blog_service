package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohuazjg/blog_service/pkg/app"
	"github.com/xiaohuazjg/blog_service/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				reponse := app.NewResponse(c)
				reponse.ToErrorResponse(errcode.TooManyRequest)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
