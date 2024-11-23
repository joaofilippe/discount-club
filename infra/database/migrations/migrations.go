package migrations

import (
	"github.com/joaofilippe/discount-club/infra/database"
	discountmigrations "github.com/joaofilippe/discount-club/infra/database/migrations/discount"
)

// CreateTables creates the necessary tables in the database.
func CreateTables(conn *database.Connection) error {
	err := discountmigrations.CreateTable(conn)
	return err
}
