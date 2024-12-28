package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/joaofilippe/discount-club/common/logger"
	appconfig "github.com/joaofilippe/discount-club/config"
)

var connection *Connection

// Connection holds the database connection details.
type Connection struct {
	db *sqlx.DB
}

// New creates a new database connection.
func New(log *logger.Logger, appConfig *appconfig.App) *Connection {
	if connection != nil {
		return connection
	}

	db, err := sqlx.Open("postgres", appConfig.Dsn)
	if err != nil {
		panic(fmt.Errorf("can't open connection to database: %w", err))
	}

	connection = &Connection{
		db,
	}

	return connection
}

// Get returns the database connection.
func (c *Connection) Get() *sqlx.DB {
	return c.db
}

func (c *Connection) RunMigrations() {
	driver, err := postgres.WithInstance(c.db.DB, &postgres.Config{})
	if err != nil {
		println(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://infra/database/migrations",
		"discount_club",
		driver,
	)
	if err != nil {
		println(err.Error())
		return
	}
	_ = m.Up()
}
