package irepositories

import "github.com/joaofilippe/discount-club/app/domain/entities"

type Discount interface {
	Save(discount *entities.Discount) error
}
