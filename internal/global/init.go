package global

import (
	"github.com/Orwell-coder/ginf/pkg/config"
	"github.com/Orwell-coder/ginf/pkg/database"
	"github.com/Orwell-coder/ginf/pkg/logger"
)

func Init() {
	var err error
	// 初始化配置
	Cfg, err = config.Init()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	Logger, err := logger.NewLogger(Cfg)
	//Logger, err := logger.NewLogger(Cfg)
	if err != nil {
		panic(err)
	}
	// 初始化数据库
	DB, err = database.NewMySQL(Cfg.MySQL, Logger)

	if err != nil {
		Logger.Sugar().Infof("init mysql failed: %s", err.Error())
	}
}
