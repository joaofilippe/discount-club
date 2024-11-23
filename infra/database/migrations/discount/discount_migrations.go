package discountmigrations

import (
	"github.com/joaofilippe/discount-club/infra/database"
)

func CreateDiscountTable(conn *database.DBConnection) error {
	_, err := conn.DBConnection.Exec(CreateTable)
	return err
}
