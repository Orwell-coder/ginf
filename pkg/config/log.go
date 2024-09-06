package config

import (
	"fmt"

	"github.com/Orwell-coder/ginf/pkg/util"
)

type Log struct {
	Development DevelopmentConfig `mapstructure:"development" yaml:"development" json:"development"`
	Production  ProductionConfig  `mapstructure:"production" yaml:"production" json:"production"`
}

type DevelopmentConfig struct {
	Level            string        `mapstructure:"level" yaml:"level" json:"level"`
	Encoding         string        `mapstructure:"encoding" yaml:"encoding" json:"encoding"`
	OutputPaths      []string      `mapstructure:"output_paths" yaml:"output_paths" json:"output_paths"`
	ErrorOutputPaths []string      `mapstructure:"error_output_paths" yaml:"error_output_paths" json:"error_output_paths"`
	EncoderConfig    EncoderConfig `mapstructure:"encoder_config" yaml:"encoder_config" json:"encoder_config"`
}

type ProductionConfig struct {
	Level            string           `mapstructure:"level" yaml:"level" json:"level"`
	Encoding         string           `mapstructure:"encoding" yaml:"encoding" json:"encoding"`
	OutputPaths      []string         `mapstructure:"output_paths" yaml:"output_paths" json:"output_paths"`
	ErrorOutputPaths []string         `mapstructure:"error_output_paths" yaml:"error_output_paths" json:"error_output_paths"`
	EncoderConfig    EncoderConfig    `mapstructure:"encoder_config" yaml:"encoder_config" json:"encoder_config"`
	FileOutput       FileOutputConfig `mapstructure:"file_output" yaml:"file_output" json:"file_output"`
	ErrorFileOutput  FileOutputConfig `mapstructure:"error_file_output" yaml:"error_file_output" json:"error_file_output"`
}

type EncoderConfig struct {
	TimeKey         string `mapstructure:"time_key" yaml:"time_key" json:"time_key"`
	LevelKey        string `mapstructure:"level_key" yaml:"level_key" json:"level_key"`
	NameKey         string `mapstructure:"name_key" yaml:"name_key" json:"name_key"`
	CallerKey       string `mapstructure:"caller_key" yaml:"caller_key" json:"caller_key"`
	MessageKey      string `mapstructure:"message_key" yaml:"message_key" json:"message_key"`
	StacktraceKey   string `mapstructure:"stacktrace_key" yaml:"stacktrace_key" json:"stacktrace_key"`
	LevelEncoder    string `mapstructure:"level_encoder" yaml:"level_encoder" json:"level_encoder"`
	TimeEncoder     string `mapstructure:"time_encoder" yaml:"time_encoder" json:"time_encoder"`
	DurationEncoder string `mapstructure:"duration_encoder" yaml:"duration_encoder" json:"duration_encoder"`
	CallerEncoder   string `mapstructure:"caller_encoder" yaml:"caller_encoder" json:"caller_encoder"`
}

type FileOutputConfig struct {
	Filename   string `mapstructure:"filename" yaml:"filename" json:"filename"`
	MaxSize    int    `mapstructure:"max_size" yaml:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" yaml:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups" json:"max_backups"`
	Compress   bool   `mapstructure:"compress" yaml:"compress" json:"compress"`
}

func InitLog() (*Log, error) {
	filename := "./config/log.yaml"
	config, err := util.ReadYaml(filename, Log{})

	if err != nil {
		return nil, fmt.Errorf("read log config failed. %w", err)
	}

	return config, nil
}
