package main

import (
	"time"

	"github.com/caarlos0/env/v6"
	zlog "github.com/rs/zerolog/log"
)

type Config struct {
	DB DB
}

type DB struct {
	PostgresDSN string `env:"DB_POSTGRES_DSN,required" envDefault:"host=localhost port=5432 user=user password=password dbname=wikigraph sslmode=disable"`

	MaxOpenConnections    int           `env:"DB_MAX_OPEN_CONNECTIONS" envDefault:"10"`
	MaxIdleConnections    int           `env:"DB_MAX_IDLE_CONNECTIONS" envDefault:"5"`
	MaxConnectionLifetime time.Duration `env:"DB_MAX_CONNECTION_LIFETIME" envDefault:"5m"`
}

func ReadConfig() Config {
	var conf Config
	err := env.Parse(&conf)
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to read the config")
	}

	return conf
}
