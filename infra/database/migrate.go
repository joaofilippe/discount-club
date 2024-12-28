package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func up(c *sqlx.DB) error {
	migrate, err := buildMigrate(c)
	if err != nil {
		return err
	}

	return migrate.Up()
}

func down(c *sqlx.DB) error {
	migrate, err := buildMigrate(c)
	if err != nil {
		return err
	}

	return migrate.Down()
}

func buildMigrate(c *sqlx.DB) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(c.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(
		"file://infra/database/migrations",
		"discount_club",
		driver,
	)
}
