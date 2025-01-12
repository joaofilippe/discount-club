package services

import (
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
	discountusecases "github.com/joaofilippe/discount-club/app/domain/usecases/discount"
)

type DiscountService struct {
	create *discountusecases.CreateUseCase
}

func NewDiscountService(discountRepo irepositories.Discount) (*DiscountService, error) {
	create, err := discountusecases.NewCreateUseCase(discountRepo)
	if err != nil {
		return nil, err
	}

	return &DiscountService{
		create: create,
	}, nil
}

func (ds *DiscountService) Create(discount *entities.Discount) (*entities.Discount, error) {
	return ds.create.Execute(discount)
}
