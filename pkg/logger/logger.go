package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const DefaultLogLevel = "info"

type Config struct {
	LogLevel string `json:"log_level"`
}

func NewLogger(cfg *Config) (*zerolog.Logger, error) {
	lvl, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed parse log level")

		return nil, err
	}
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).Level(lvl).With().Caller().Timestamp().Logger()

	return &logger, nil
}
