package migrations

import (
	"github.com/joaofilippe/discount-club/infra/database"
	discountmigrations "github.com/joaofilippe/discount-club/infra/database/migrations/discount"
	usersmigrations "github.com/joaofilippe/discount-club/infra/database/migrations/users"
)

// MigrateUp creates the necessary tables in the database.
func MigrateUp(conn *database.Connection) error {

	err := usersmigrations.CreateTable(conn)
	err = discountmigrations.CreateTable(conn)
	return err
}
