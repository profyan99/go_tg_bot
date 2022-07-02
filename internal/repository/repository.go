package repository

import (
	entSql "entgo.io/ent/dialect/sql"
	"github.com/rs/zerolog"
	"go-tg-bot/internal/repository/ent"
)

type Repository struct {
	Client *ent.Client
	log    *zerolog.Logger
}

func NewRepository(db *ent.Client, logger *zerolog.Logger) *Repository {
	return &Repository{
		Client: db,
		log:    logger,
	}
}

func NewDBClient(drv *entSql.Driver) (client *ent.Client, err error) {
	client = ent.NewClient(ent.Driver(drv))
	return client, nil
}
