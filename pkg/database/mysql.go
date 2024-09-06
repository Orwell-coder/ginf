package database

import (
	"fmt"
	"time"

	"github.com/Orwell-coder/ginf/pkg/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMySQL(cfg *config.MySQL, zapLogger *zap.Logger) (*gorm.DB, error) {
	// 创建一个新的gorm logger
	gormLogger := logger.New(
		NewZapAdapter(zapLogger), // 使用zap logger作为底层logger
		logger.Config{
			SlowThreshold:             time.Duration(cfg.SlowQueryThreshold) * time.Millisecond,
			LogLevel:                  getLogLevel(cfg.LogModeLevel),
			IgnoreRecordNotFoundError: cfg.IgnoreRecordNotFoundError,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Params)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger:                                   gormLogger,
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键检测
		NamingStrategy: schema.NamingStrategy{
			SingularTable: cfg.SingularTable,
			TablePrefix:   cfg.TablePrefix,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func getLogLevel(level string) logger.LogLevel {
	switch level {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}

// ZapAdapter 适配zap logger到gorm logger接口
type ZapAdapter struct {
	ZapLogger *zap.Logger
}

func NewZapAdapter(zapLogger *zap.Logger) *ZapAdapter {
	return &ZapAdapter{ZapLogger: zapLogger}
}

func (a *ZapAdapter) Printf(format string, args ...interface{}) {
	a.ZapLogger.Sugar().Infof(format, args...)
}
