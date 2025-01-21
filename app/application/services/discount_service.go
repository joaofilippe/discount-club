package services

import (
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
	discountusecases "github.com/joaofilippe/discount-club/app/domain/usecases/discount"
)

type DiscountService struct {
	create  *discountusecases.CreateUseCase
	getByID *discountusecases.GetByIDUseCase
}

func NewDiscountService(discountRepo irepositories.Discount) (*DiscountService, error) {
	create, err := discountusecases.NewCreateUseCase(discountRepo)
	if err != nil {
		return nil, err
	}

	getByID, err := discountusecases.NewGetByIDUseCase(discountRepo)
	if err != nil {
		return nil, err
	}

	return &DiscountService{
		create,
		getByID,
	}, nil
}

func (ds *DiscountService) Create(discount *entities.Discount) (*entities.Discount, error) {
	return ds.create.Execute(discount)
}

func (ds *DiscountService) GetByID(id string) (*entities.Discount, error) {
	return ds.getByID.Execute(id)
}
