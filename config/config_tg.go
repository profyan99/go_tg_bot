package config

import "github.com/spf13/viper"

const (
	envDebug      = "TG_DEBUG"
	envTgApiKey   = "TG_API_KEY"
	defaultDebug  = false
	defaultApiKey = ""
)

type TgConfig struct {
	Debug  bool   `json:"debug"`
	ApiKey string `json:"api_key"`
}

func NewTgConfig(v *viper.Viper) *TgConfig {
	v.SetDefault(envDebug, defaultDebug)
	v.SetDefault(envTgApiKey, defaultApiKey)
	return &TgConfig{
		Debug:  v.GetBool(envDebug),
		ApiKey: v.GetString(envTgApiKey),
	}
}
