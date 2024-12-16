package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //necessary to config

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
