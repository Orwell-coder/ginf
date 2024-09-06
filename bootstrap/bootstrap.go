package bootstrap

import (
	"github.com/Orwell-coder/ginf/internal/global"
	"github.com/gin-gonic/gin"
)

func init() {
	global.Init()
	InitGinConfig()
}

func InitGinConfig() {
	if global.Cfg.Env {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
