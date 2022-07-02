package db

import (
	"database/sql"

	entSql "entgo.io/ent/dialect/sql"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"

	// Need to work with migration files.
	_ "github.com/lib/pq"
)

func GetDriver(conf *Config, log *zerolog.Logger) (*entSql.Driver, error) {
	db, err := GetConnect(conf, log)
	if err != nil {
		return nil, err
	}
	return entSql.OpenDB(conf.DriverName, db), nil
}

func Migrations(conf *Config, log *zerolog.Logger) error {
	db, err := sql.Open(conf.DriverName, conf.DSN())
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetBaseFS(nil)

	if err := goose.Up(db, conf.MigrationDirectory); err != nil {
		return err
	}

	return nil
}
