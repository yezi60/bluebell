package middlewares

import (
	"net/http"
	"time"

	"github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
)

// ReteLimitMiddleware 填充速率，总容量
func ReteLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)

	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		// 如果剩余的令牌数量为0，就返回
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		// 取到令牌就放行
		c.Next()
	}

}
