package discountusecases_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	errormessages "github.com/joaofilippe/discount-club/common/errors/messages"
	"github.com/joaofilippe/discount-club/domain/entities"
	discountusecases "github.com/joaofilippe/discount-club/domain/usecases/discount"
)

type DiscountRepoMock struct {
	mock.Mock
}

func (d *DiscountRepoMock) Save(discount *entities.Discount) error {
	if discount == nil {
		return errors.New(errormessages.ErrNilDiscount)
	}

	return nil
}

var _ = Describe("Create.Usecase", func() {
	var (
		discountRepo    *DiscountRepoMock
		discountUseCase *discountusecases.CreateUseCase
		err error
	)

	Describe("New CreateUseCase", func() {
		Context("whith a valid discount repository", func() {
			BeforeEach(func() {
				discountRepo = &DiscountRepoMock{}
				discountUseCase, err = discountusecases.NewCreateUseCase(discountRepo)
			})

			It("should return a valid use case", func() {
				Expect(discountUseCase).NotTo(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
