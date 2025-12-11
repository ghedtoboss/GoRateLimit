package limiter

import (
	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(rlm *RateLimitManager) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		bucket := rlm.GetOrCreate(c.ClientIP())
		if !bucket.Allow() {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too Many Requests"})
			return
		}
		c.Next()
	})
}