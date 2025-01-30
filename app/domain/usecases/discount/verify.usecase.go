package discountusecases

import (
	discounterrors "github.com/joaofilippe/discount-club/app/common/myerrors/discount"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
	"strings"
)

type VerifyUsecase struct {
	discountRep irepositories.Discount
}

func NewVerifyUsecase(discountRep irepositories.Discount) (*VerifyUsecase, error) {
	if discountRep == nil {
		return nil, discounterrors.ErrNilDiscountRepo
	}

	return &VerifyUsecase{
		discountRep: discountRep,
	}, nil
}

func (vu *VerifyUsecase) Execute(code string) (bool, error) {
	discount, err := vu.discountRep.GetByCode(code)
	if err != nil {
		return false, err
	}

	if discount == nil {
		return false, discounterrors.ErrDiscountNotFound
	}

	if strings.EqualFold(discount.Code(), code) {
		return true, nil
	}

	return false, nil
}
