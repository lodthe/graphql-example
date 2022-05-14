package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/lodthe/graphql-example/internal/gqlgenerated"
	"github.com/lodthe/graphql-example/internal/match"
	"github.com/lodthe/graphql-example/internal/resolve"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	conf := ReadConfig()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	_, cancel := context.WithCancel(context.Background())
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	db, err := setupDatabaseConnection(conf.DB)
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to setup database connection")
	}
	defer db.Close()

	repo := match.NewRepository(db)

	http.Handle("/", playground.Handler("Demo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(gqlgenerated.NewExecutableSchema(
		gqlgenerated.Config{
			Resolvers: resolve.NewResolver(repo),
		},
	)))

	go func() {
		err := http.ListenAndServe(conf.ServerAddress, nil)
		if err != nil {
			zlog.Fatal().Err(err).Msg("server failed")
		}
	}()

	zlog.Info().Str("address", conf.ServerAddress).Msg("server has been started")

	<-stop
	cancel()
}

func setupDatabaseConnection(config DB) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", config.PostgresDSN)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(config.MaxConnectionLifetime)
	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)

	return db, nil
}
