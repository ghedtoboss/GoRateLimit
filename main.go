package main

import (
	"GoRateLimit/limiter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create limit manager
	rlm := limiter.NewRateLimitManager(5, 5)

	// gin router
	r := gin.Default()

	// use middleware
	r.Use(limiter.RateLimitMiddleware(rlm))

	// test endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// run server
	r.Run(":8080")
}
