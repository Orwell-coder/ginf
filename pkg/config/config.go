package config

import (
	"fmt"

	"github.com/Orwell-coder/ginf/pkg/util"
)

type Config struct {
	Log   *Log   `yaml:"log" json:"log"`
	MySQL *MySQL `yaml:"mysql" json:"mysql"`
	Redis *Redis `yaml:"redis" json:"redis"`
	Env   bool   `yaml:"env" json:"env"`
	Host  string `yaml:"host" json:"host"`
	Port  int    `yaml:"port" json:"port"`
}

func InitConfig() (*Config, error) {
	filename := "./config/Config.yaml"
	config, err := util.ReadYaml(filename, Config{})

	if err != nil {
		return nil, fmt.Errorf("read Config config failed. %w", err)
	}

	return config, nil
}
