package repositories

import (
	dbmodels "github.com/joaofilippe/discount-club/app/application/repositories/models"
	"github.com/joaofilippe/discount-club/app/application/repositories/queries"
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/infra/database"
)

type DiscountRepo struct {
	conn *database.Connection
}

func NewDiscountRepo(conn *database.Connection) *DiscountRepo {
	return &DiscountRepo{
		conn: conn,
	}
}

func (dr *DiscountRepo) Save(discount *entities.Discount) error {
	tx, err := dr.conn.Get().Beginx()
	if err != nil {
		return err
	}

	discountDB := new(dbmodels.Discount)
	discountDB.FromEntity(discount)

	_, err = tx.NamedExec(queries.InsertDiscount, discountDB)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
