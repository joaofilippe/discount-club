package iservices

import "github.com/joaofilippe/discount-club/app/domain/entities"

type IDiscount interface {
	Create(discount *entities.Discount) (*entities.Discount, error)
	GetByID(id string) (*entities.Discount, error)
	Verify(code string) (bool, error)
}
