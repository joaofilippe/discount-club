package discountmigrations

import (
	"github.com/joaofilippe/discount-club/infra/database"
)

// CreateTable creates the discount table in the database.
func CreateTable(conn *database.Connection) error {
	tx := conn.Get().MustBegin()
	_, err := tx.Exec(createTableQuery)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// DropTable drops the discount table from the database.
func DropTable(conn *database.Connection) error {
	tx := conn.Get().MustBegin()
	_, err := tx.Exec(dropTableQuery)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
