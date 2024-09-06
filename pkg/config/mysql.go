package config

import (
	"fmt"

	"github.com/Orwell-coder/ginf/pkg/util"
)

type MySQL struct {
	Host                      string `yaml:"host" json:"host"`
	Port                      int    `yaml:"port" json:"port"`
	User                      string `yaml:"user" json:"user"`
	Password                  string `yaml:"password" json:"password"`
	DbName                    string `yaml:"db_name" json:"db_name" mapstructtag:"db_name"`
	Params                    string `yaml:"params" json:"params"`
	MaxIdleConns              int    `yaml:"max_idle_conns" json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns              int    `yaml:"max_open_conns" json:"max_open_conns" mapstructure:"max_open_conns"`
	SingularTable             bool   `yaml:"singular_table" json:"singular_table" mapstructure:"singular_table"`
	TablePrefix               string `yaml:"table_prefix" json:"table_prefix" mapstructure:"table_prefix"`
	LogMode                   bool   `yaml:"log_mode" json:"log_mode" mapstructure:"log_mode"`
	LogModeLevel              string `yaml:"log_mode_level" json:"log_mode_level" mapstructure:"log_mode_level"`
	LogZap                    bool   `yaml:"log_zap" json:"log_zap" mapstructure:"log_zap"`
	LogZapLevel               string `yaml:"log_zap_level" json:"log_zap_level" mapstructure:"log_zap_level"`
	SlowQueryThreshold        int    `yaml:"slow_query_threshold" json:"slow_query_threshold" mapstructure:"slow_query_threshold"`
	IgnoreRecordNotFoundError bool   `yaml:"ignore_record_not_found_error" json:"ignore_record_not_found_error" mapstructure:"ignore_record_not_found_error"`
}

func InitMySQL() (*MySQL, error) {
	filename := "./config/mysql.yaml"
	config, err := util.ReadYaml(filename, MySQL{})

	if err != nil {
		return nil, fmt.Errorf("read mysql config failed. %w", err)
	}

	return config, nil
}
