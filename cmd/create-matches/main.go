package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/lodthe/graphql-example/internal/match"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	conf := ReadConfig()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := setupDatabaseConnection(conf.DB)
	if err != nil {
		zlog.Fatal().Err(err).Msg("failed to setup database connection")
	}
	defer db.Close()

	repo := match.NewRepository(db)

	for i := 0; i < 15; i++ {
		finished := true
		if rand.Intn(3) == 0 {
			finished = false
		}

		m := &match.Match{
			ID: uuid.New(),
			State: match.State{
				CreatedAt:  time.Now().Add(time.Duration(rand.Intn(10000)) * time.Second),
				IsFinished: finished,
				Comments:   nil,
				Scoreboard: match.Scoreboard{},
			},
		}
		for j := 0; j < 4+rand.Intn(3); j++ {
			alive := true
			if rand.Intn(3) == 0 {
				alive = false
			}

			role := match.RoleVillager
			if rand.Intn(3) == 0 {
				role = match.RoleMafia
			}

			names := []string{"Igor", "John", "Sabina", "Alice", "Polina", "DarkLord", "VVV", "Desant", "Ivan", "Bill", "Elizabeth", "Michael", "Margo", "Suzanna"}
			surnames := []string{"Black", "Green", "Widow", "Polime", "Sally", "Saddlton", "Diavol"}
			m.State.Scoreboard.Players = append(m.State.Scoreboard.Players, match.Player{
				Username: names[rand.Intn(len(names))] + " " + surnames[rand.Intn(len(surnames))],
				Role:     role,
				IsAlive:  alive,
				Kills:    rand.Intn(2),
			})
		}

		err := repo.Create(m)
		if err != nil {
			zlog.Fatal().Err(err).Interface("match", m).Msg("failed to create")
		}

		zlog.Info().Msg("created")
	}
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
