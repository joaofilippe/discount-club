package repositories

import (
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/infra/database"
)

type DiscountRepo struct {
	conn *database.Connection
}

func (dr *DiscountRepo) Save(discount *entities.Discount) error {
	return nil
}
