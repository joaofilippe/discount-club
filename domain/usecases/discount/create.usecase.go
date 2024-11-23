package discountusecases

import (
	"math/rand/v2"
	"time"

	"github.com/joaofilippe/discount-club/domain/entities"
	"github.com/joaofilippe/discount-club/domain/irepositories"
)

type CreateUseCase struct {
	DiscountRepo irepositories.Discount
}

func NewCreateUseCase(discountRepo irepositories.Discount) *CreateUseCase {
	return &CreateUseCase{
		DiscountRepo: discountRepo,
	}
}

func (uc *CreateUseCase) Execute(discount *entities.Discount) error {
	
	
	return nil
}



func generateCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	
	pcg := rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Second()))
	r := rand.New(pcg)

	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[r.IntN(len(charset))]
	}

	return ""
}
