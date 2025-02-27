package discountusecases

import (
	"fmt"
	"math/rand/v2"
	"time"

	discounterrors "github.com/joaofilippe/discount-club/app/common/myerrors/discount"
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
)

// CreateUseCase handles the creation of discounts
type CreateUseCase struct {
	DiscountRepo irepositories.Discount
	randSource   *rand.Rand
}

// NewCreateUseCase creates a new instance of CreateUseCase
func NewCreateUseCase(discountRepo irepositories.Discount) (*CreateUseCase, error) {
	if discountRepo == nil {
		return nil, discounterrors.ErrNilDiscountRepo
	}

	pcg := rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Second()))
	r := rand.New(pcg)

	return &CreateUseCase{
		DiscountRepo: discountRepo,
		randSource:   r,
	}, nil
}

// Execute processes the discount creation request
func (uc *CreateUseCase) Execute(request *entities.Discount) (*entities.Discount, error) {
	if request == nil {
		return &entities.Discount{}, discounterrors.ErrNoDiscountProvided
	}

	if err := uc.validatedRequest(request); err != nil {
		return &entities.Discount{}, err
	}

	initializedDiscount := uc.prepareDiscount(request)

	if err := uc.DiscountRepo.Save(initializedDiscount); err != nil {
		return &entities.Discount{}, err
	}

	return initializedDiscount, nil
}

// prepareDiscount initializes a new discount with a generated code
func (uc *CreateUseCase) prepareDiscount(discount *entities.Discount) *entities.Discount {
	now := time.Now()
	code := uc.generateCode(uc.randSource)
	return discount.CopyWith(
		nil,
		nil,
		nil,
		&code,
		nil,
		nil,
		nil,
		nil,
		nil,
		&now,
		nil,
		nil,
		nil,
	)

}

// generateCode generates a random discount code
func (uc *CreateUseCase) generateCode(r *rand.Rand) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[r.IntN(len(charset))]
	}

	return string(b)
}

// validatedRequest validates the discount request
func (uc *CreateUseCase) validatedRequest(discount *entities.Discount) error {
	if discount.Code() != "" {
		return discounterrors.ErrCodeProvidedOnCreate
	}

	if discount.Percentage() <= 0 || discount.Percentage() > 100 {
		return fmt.Errorf("invalid percentage: %d", discount.Percentage())
	}

	if discount.Times() <= 0 {
		return fmt.Errorf("invalid times: %d", discount.Times())
	}

	if discount.Description() == "" {
		return fmt.Errorf("description is required")
	}

	return nil
}
