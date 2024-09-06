package config

import (
	"fmt"

	"github.com/Orwell-coder/ginf/pkg/util"
)

type Redis struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
}

func InitRedis() (*Redis, error) {
	filename := "./config/redis.yaml"
	config, err := util.ReadYaml(filename, Redis{})

	if err != nil {
		return nil, fmt.Errorf("read redis config failed. %w", err)
	}

	return config, nil
}
