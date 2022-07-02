package config

import (
	"github.com/spf13/viper"
	"go-tg-bot/pkg/db"
	"go-tg-bot/pkg/logger"
)

const (
	defaultValue = ""
	envPrefix    = "SLACK_BOT_BIRTHDAY"
)

type Config struct {
	Viper  *viper.Viper   `json:"viper"`
	Logger *logger.Config `json:"logger"`
	DB     *db.Config     `json:"db"`
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
	}
}
