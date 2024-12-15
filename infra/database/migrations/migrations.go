package migrations

import (
	"github.com/joaofilippe/discount-club/infra/database"
	discountmigrations "github.com/joaofilippe/discount-club/infra/database/migrations/discount"
	objecttypesmigrations "github.com/joaofilippe/discount-club/infra/database/migrations/object_types"
	usersmigrations "github.com/joaofilippe/discount-club/infra/database/migrations/users"
)

// Up creates the necessary tables in the database.
func Up(conn *database.Connection) error {
	err := objecttypesmigrations.CreateObjectTypeENUM(conn)
	if err != nil {
		return err
	}

	err = usersmigrations.CreateTable(conn)
	if err != nil {
		return err
	}

	err = discountmigrations.CreateTable(conn)
	if err != nil {
		return err
	}

	return err
}
func Down(conn *database.Connection) error {
	err := objecttypesmigrations.DropObjectTypeENUM(conn)
	if err != nil {
		return err
	}

	err = discountmigrations.DropTable(conn)
	if err != nil {
		return err
	}

	return usersmigrations.DropTable(conn)

}
