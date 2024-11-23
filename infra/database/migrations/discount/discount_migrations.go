package discountmigrations

import (
	"github.com/joaofilippe/discount-club/infra/database"
)


// CreateDiscountTable creates the discount table in the database.
func CreateDiscountTable(conn *database.Connection) error {
	tx := conn.Get().MustBegin()
	_, err := tx.Exec(CreateTable)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
