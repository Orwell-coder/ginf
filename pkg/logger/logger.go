package logger

import (
	"os"

	"github.com/Orwell-coder/ginf/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	var zapConfig zap.Config
	var writers []zapcore.WriteSyncer

	if !cfg.Env {
		zapConfig = zap.Config{
			Level:             zap.NewAtomicLevelAt(getLogLevel(cfg.Log.Production.Level)),
			Encoding:          cfg.Log.Production.Encoding,
			EncoderConfig:     getEncoderConfig(cfg.Log.Production.EncoderConfig),
			OutputPaths:       cfg.Log.Production.OutputPaths,
			ErrorOutputPaths:  cfg.Log.Production.ErrorOutputPaths,
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
		}

		// 添加文件输出与轮转
		if contains(cfg.Log.Production.OutputPaths, "file") {
			fileWriter := zapcore.AddSync(&lumberjack.Logger{
				Filename:   cfg.Log.Production.FileOutput.Filename,
				MaxSize:    cfg.Log.Production.FileOutput.MaxSize,
				MaxBackups: cfg.Log.Production.FileOutput.MaxBackups,
				MaxAge:     cfg.Log.Production.FileOutput.MaxAge,
				Compress:   cfg.Log.Production.FileOutput.Compress,
			})
			writers = append(writers, fileWriter)
		}
	} else {
		zapConfig = zap.Config{
			Level:             zap.NewAtomicLevelAt(getLogLevel(cfg.Log.Development.Level)),
			Encoding:          cfg.Log.Development.Encoding,
			EncoderConfig:     getEncoderConfig(cfg.Log.Development.EncoderConfig),
			OutputPaths:       cfg.Log.Development.OutputPaths,
			ErrorOutputPaths:  cfg.Log.Development.ErrorOutputPaths,
			Development:       true,
			DisableCaller:     false,
			DisableStacktrace: false,
		}
	}

	// 添加控制台输出
	if contains(zapConfig.OutputPaths, "stdout") {
		writers = append(writers, zapcore.AddSync(os.Stdout))
	}

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig.EncoderConfig),
		zapcore.NewMultiWriteSyncer(writers...),
		zapConfig.Level,
	)

	// 创建logger
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger, nil
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// ... (getLogLevel 和 getEncoderConfig 函数保持不变)

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getEncoderConfig(cfg config.EncoderConfig) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        cfg.TimeKey,
		LevelKey:       cfg.LevelKey,
		NameKey:        cfg.NameKey,
		CallerKey:      cfg.CallerKey,
		MessageKey:     cfg.MessageKey,
		StacktraceKey:  cfg.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
