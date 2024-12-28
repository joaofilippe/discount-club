package database

import (
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/joaofilippe/discount-club/common/logger"
	appconfig "github.com/joaofilippe/discount-club/config"
)

var connection *Connection

// Connection holds the database connection details.
type Connection struct {
	app *appconfig.App
	db  *sqlx.DB
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
		appConfig,
		db,
	}

	return connection
}

// Get returns the database connection.
func (c *Connection) Get() *sqlx.DB {
	return c.db
}

func (c *Connection) RunMigrations() {
	if os.Getenv("RESET_DB") == "true" {
		if err := down(c.db); err != nil {
			c.app.Logger.Fatalf(err)
		}
	}

	if err := up(c.db); err != nil && err.Error() != "no change" {
		c.app.Logger.Fatalf(err)
	}
}
