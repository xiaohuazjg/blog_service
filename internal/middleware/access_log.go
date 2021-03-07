package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/xiaohuazjg/blog_service/global"
	"github.com/xiaohuazjg/blog_service/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandleFunc {
	return func(c *gin.Context) {
		bodywriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodywriter
		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodywriter.body.String(),
		}
		s := "access_log:method:%s, status_code:%d," +
			"begin_time:%d,end_time:%d"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodywriter.Status(),
			beginTime,
			endTime,
		)

	}
}
