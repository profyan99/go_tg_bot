package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
	"go-tg-bot/config"
	"go-tg-bot/internal/app/commands"
	"go-tg-bot/internal/app/pagination"
	"go-tg-bot/internal/repository"
	"go-tg-bot/internal/service/product"
	"go-tg-bot/pkg/db"
	lc "go-tg-bot/pkg/logger"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Error().Err(err).Msg("shutting down")
		os.Exit(1)
	}
}

func run() error {
	// Configuration
	cfg := config.LoadConfig()

	bot, err := tgbotapi.NewBotAPI(cfg.TG.ApiKey)
	if err != nil {
		return err
	}

	// Logging
	logger, err := lc.NewLogger(cfg.Logger)
	if err != nil {
		return err
	}

	bot.Debug = cfg.TG.Debug

	logger.Info().Str("Authorized on account", bot.Self.UserName).Msg("TG")

	err = db.Migrations(cfg.DB, logger)
	if err != nil {
		return err
	}
	dbDriver, err := db.GetDriver(cfg.DB, logger)
	if err != nil {
		return err
	}
	dbClient, err := repository.NewDBClient(dbDriver)
	if err != nil {
		return err
	}
	defer dbClient.Close()

	// repository
	_ = repository.NewRepository(dbClient, logger)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	paginationService := pagination.NewPagination(productService)
	commander := commands.NewCommander(bot, productService, paginationService)

	for update := range updates {
		commander.HandleUpdate(update)
	}

	return nil
}
