package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config это глобальный объект, который содержит все переменные из структуры appConfig.
var Config appConfig

type appConfig struct {
	CheckConfig string
	BindIP      string
	Port        string
}

// LoadConfig подгружает конфигурацию из файлов
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("connection")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("rcms")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
