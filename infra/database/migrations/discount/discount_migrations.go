package discountmigrations

import (
	"github.com/joaofilippe/discount-club/infra/database"
)


// CreateDiscountTable creates the discount table in the database.
func CreateTable(conn *database.Connection) error {
	tx := conn.Get().MustBegin()
	_, err := tx.Exec(createTableQuery)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
