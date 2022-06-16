package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config это глобальный объект, который содержит все переменные уровня приложения.
var Config appConfig

type appConfig struct {
	// Пример переменной для теста конфига
	ConfigVar string
}

// LoadConfig подгружает конфигурацию из файлов
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("example")
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
