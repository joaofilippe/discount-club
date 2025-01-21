package discountusecases

import (
	"github.com/google/uuid"
	discounterrors "github.com/joaofilippe/discount-club/app/common/myerrors/discount"
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
)

type GetByIDUseCase struct {
	DiscountRepo irepositories.Discount
}

func NewGetByIDUseCase(discountRepo irepositories.Discount) (*GetByIDUseCase, error) {
	if discountRepo == nil {
		return nil, discounterrors.ErrNilDiscountRepo
	}

	return &GetByIDUseCase{
		DiscountRepo: discountRepo,
	}, nil
}

func (uc *GetByIDUseCase) Execute(id string) (*entities.Discount, error) {
	if id == "" {
		return entities.NewEmptyDiscount(), discounterrors.ErrNoDiscountIDProvided
	}

	uuidID, err := uuid.Parse(id)
	if err != nil || uuidID == uuid.Nil {
		return entities.NewEmptyDiscount(), discounterrors.ErrDiscountNotFound
	}

	discount, err := uc.DiscountRepo.GetByID(uuidID)
	if err != nil {
		return entities.NewEmptyDiscount(), err
	}

	if discount == nil {
		return entities.NewEmptyDiscount(), discounterrors.ErrDiscountNotFound
	}

	if discount.ID() == uuid.Nil {
		return entities.NewEmptyDiscount(), discounterrors.ErrDiscountNotFound
	}

	return discount, nil
}