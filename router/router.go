package router

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	// 路由模块化配置
	apiV1 := router.Group("/api/v1")
	{
		RegisterTestRoutes(apiV1)
	}

	return router
}
