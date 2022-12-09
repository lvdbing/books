package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lvdbing/books/blog-service/global"
	"github.com/lvdbing/books/blog-service/pkg/app"
	"github.com/lvdbing/books/blog-service/pkg/email"
	"github.com/lvdbing/books/blog-service/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	mailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		Username: global.EmailSetting.Username,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)

				err := mailer.SendEmail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出, 发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendEmail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
