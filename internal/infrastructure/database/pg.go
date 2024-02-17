package database

import (
	"context"
	"database/sql"
	"os"

	migr "github.com/lohuza/relayer/cmd/migrator/migrations"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
)

func InitPg(connectionString string) *bun.DB {
	log.Info().Msg("connecting to postgres db")
	pgDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))
	db := bun.NewDB(pgDb, pgdialect.New())
	err := db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	env := os.Getenv("APP_MODE")
	if env == "debug" {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	}

	err = InitMigrations(db)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to check migration status")
	}

	return db
}

func InitMigrations(db *bun.DB) error {
	migr := migrate.NewMigrator(db, migr.Migrations)
	err := migr.Init(context.TODO())
	if err != nil {
		return err
	}
	_, err = migr.Migrate(context.TODO())
	return err
}
