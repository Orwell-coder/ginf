package global

import (
	"github.com/Orwell-coder/ginf/pkg/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Logger *zap.Logger
	Cfg    *config.Config
)
