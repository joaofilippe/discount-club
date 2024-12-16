package discountusecases

import (
	"errors"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
	errormessages "github.com/joaofilippe/discount-club/common/errors/messages"
	"github.com/joaofilippe/discount-club/domain/entities"
	"github.com/joaofilippe/discount-club/domain/irepositories"
)

type CreateUseCase struct {
	DiscountRepo irepositories.Discount
}

func NewCreateUseCase(discountRepo irepositories.Discount) (*CreateUseCase, error) {
	if discountRepo == nil {
		return nil, errors.New(errormessages.ErrNilDiscountRepo)
	}

	return &CreateUseCase{
		DiscountRepo: discountRepo,
	}, nil
}

func (uc *CreateUseCase) Execute(discount *entities.Discount) error {
	discount.SetID(uuid.New())
	discount.SetCode(generateCode())

	return uc.DiscountRepo.Save(discount)	
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
