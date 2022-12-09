package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lvdbing/books/blog-service/global"
	"github.com/lvdbing/books/blog-service/pkg/logger"
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

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter

		begin := time.Now().Unix()
		c.Next()
		end := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		format := "access log: method: %s, status_code: %d, begin: %d, end: %d"
		global.Logger.WithFields(fields).Infof(c, format,
			c.Request.Method,
			bodyWriter.Status(),
			begin,
			end,
		)
	}
}
