package iservices

import "github.com/joaofilippe/discount-club/app/domain/entities"

type IDiscountServices interface {
	Create(discount *entities.Discount) (*entities.Discount, error)
}