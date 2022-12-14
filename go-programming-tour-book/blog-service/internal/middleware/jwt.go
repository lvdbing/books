package middleware

import (
	"github.com/lvdbing/books/blog-service/global"
	"github.com/lvdbing/books/blog-service/pkg/app"
	"github.com/lvdbing/books/blog-service/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		var ecode = errcode.Success

		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			global.Logger.Error(c, "valid token error")
			resp := app.NewResponse(c)
			resp.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
