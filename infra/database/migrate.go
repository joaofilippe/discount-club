package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Up(c *sqlx.DB) error {
	migrate, err := buildMigrate(c)
	if err != nil {
		return err
	}

	return migrate.Up()
}

func Down(c *sqlx.DB) error {
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
