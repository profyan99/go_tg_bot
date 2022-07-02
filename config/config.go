package config

import (
	"github.com/spf13/viper"
	"go-tg-bot/pkg/db"
	"go-tg-bot/pkg/logger"
)

const (
	envPrefix = "GO_TG_BOT"
)

type Config struct {
	Viper  *viper.Viper   `json:"viper"`
	Logger *logger.Config `json:"logger"`
	DB     *db.Config     `json:"db"`
	TG     *TgConfig      `json:"tg"`
}

// LoadConfig creates a Config object that is filled with values from environment variables or set default values
func LoadConfig() *Config {
	v := viper.New()
	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	return &Config{
		Viper:  v,
		Logger: NewLoggerConfig(v),
		DB:     NewDBConfig(v),
		TG:     NewTgConfig(v),
	}
}
