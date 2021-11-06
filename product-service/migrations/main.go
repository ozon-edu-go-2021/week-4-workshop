package main

import (
	"database/sql"
	"embed"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ozonmp/week-4-workshop/product-service/internal/config"
	"github.com/pressly/goose/v3"

	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	db, err := sql.Open("pgx", cfg.DB.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed SQL open")
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	const cmd = "up"
	err = goose.Run(cmd, db, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("Goose error")
	}
}