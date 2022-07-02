package config

import (
	"github.com/spf13/viper"
	"go-tg-bot/pkg/db"
)

const (
	envDBHost               = "DB_HOST"
	envDBPort               = "DB_PORT"
	envDBName               = "DB_NAME"
	envDBUser               = "DB_USER"
	envDBPassword           = "DB_PASSWORD"
	envDBSchema             = "DB_SCHEMA"
	envDBSslMode            = "DB_SSL_MODE"
	envDBMigrationDirectory = "DB_MIGRATION_DIRECTORY"
	envDBDriverName         = "DB_DRIVER_NAME"
)

func NewDBConfig(v *viper.Viper) *db.Config {
	v.SetDefault(envDBHost, db.DefaultHost)
	v.SetDefault(envDBPort, db.DefaultPort)
	v.SetDefault(envDBName, db.DefaultDatabaseName)
	v.SetDefault(envDBUser, db.DefaultUser)
	v.SetDefault(envDBSslMode, db.DefaultSSLMode)
	v.SetDefault(envDBMigrationDirectory, db.DefaultMigrationDirectory)
	v.SetDefault(envDBDriverName, db.DefaultDriverName)
	return &db.Config{
		Host:               v.GetString(envDBHost),
		Port:               v.GetString(envDBPort),
		DatabaseName:       v.GetString(envDBName),
		User:               v.GetString(envDBUser),
		Password:           v.GetString(envDBPassword),
		Schema:             v.GetString(envDBSchema),
		SslMode:            v.GetString(envDBSslMode),
		MigrationDirectory: v.GetString(envDBMigrationDirectory),
		DriverName:         v.GetString(envDBDriverName),
	}
}
