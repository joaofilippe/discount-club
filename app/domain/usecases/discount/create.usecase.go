package discountusecases

import (
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
	discounterrors "github.com/joaofilippe/discount-club/app/common/myerrors/discount"
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
)

type CreateUseCase struct {
	DiscountRepo irepositories.Discount
}

func NewCreateUseCase(discountRepo irepositories.Discount) (*CreateUseCase, error) {
	if discountRepo == nil {
		return nil, discounterrors.ErrNilDiscountRepo
	}

	return &CreateUseCase{
		DiscountRepo: discountRepo,
	}, nil
}

func (uc *CreateUseCase) Execute(discount *entities.Discount) (*entities.Discount, error) {
	if discount == nil {
		return nil, discounterrors.ErrNoDiscountProvided
	}

	if discount.Code() != "" {
		return nil, discounterrors.ErrCodeProvidedOnCreate
	}

	discount.SetID(uuid.New())
	discount.SetCode(generateCode())

	return discount, uc.DiscountRepo.Save(discount)
}

func generateCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	pcg := rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Second()))
	r := rand.New(pcg)

	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[r.IntN(len(charset))]
	}

	return string(b)
}
