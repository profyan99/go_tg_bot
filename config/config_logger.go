package config

import (
	"github.com/spf13/viper"
	"go-tg-bot/pkg/logger"
)

const (
	envLogLevel = "LOG_LEVEL"
)

func NewLoggerConfig(v *viper.Viper) *logger.Config {
	v.SetDefault(envLogLevel, logger.DefaultLogLevel)
	return &logger.Config{
		LogLevel: v.GetString(envLogLevel),
	}
}
