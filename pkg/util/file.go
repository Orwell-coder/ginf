package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadYaml[T any](path string, obj T) (*T, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	// fd, err := os.Open(path)

	// if err != nil {
	// 	return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	// }
	// defer fd.Close()
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}
	v.Unmarshal(&obj)
	return &obj, nil

}
