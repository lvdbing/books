package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lvdbing/books/blog-service/pkg/app"
	"github.com/lvdbing/books/blog-service/pkg/errcode"
	"github.com/lvdbing/books/blog-service/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				resp := app.NewResponse(c)
				resp.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
