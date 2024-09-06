package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterTestRoutes(r *gin.RouterGroup) {
	group := r.Group("/test")
	{
		group.GET("/index", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

	}
}
