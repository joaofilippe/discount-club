package discountusecases_test

import (
	"errors"
	"time"

	"github.com/google/uuid"
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
		discount        *entities.Discount
		err             error
	)

	Describe("NewCreateUseCase", func() {
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

		Context("whith a nil discount repository", func() {
			BeforeEach(func() {
				discountUseCase, err = discountusecases.NewCreateUseCase(nil)
			})

			It("should return an error", func() {
				Expect(discountUseCase).To(BeNil())
				Expect(err).To(Equal(errors.New(errormessages.ErrNilDiscountRepo)))
			})
		})
	})

	Describe("Execute", func() {
		BeforeEach(func() {
			discountRepo = &DiscountRepoMock{}
			discountUseCase, _ = discountusecases.NewCreateUseCase(discountRepo)

			startTime, _ := time.Parse("2006-01-02", "2021-01-31")
			endTime := startTime.Add(time.Hour * 24 * 31)
			discount = entities.NewDiscount(
				uuid.New(),
				uuid.New(),
				10,
				startTime,
				endTime,
				1,
				"test discount",
			)
		})

		Context("whith a valid discount", func() {
			BeforeEach(func() {
				discountRepo.On("Save", discount).Return(nil)
				discount, err = discountUseCase.Execute(discount)
			})

			It("should return no error", func() {
				Expect(err).To(BeNil())
				Expect(discount).NotTo(BeNil())
				Expect(discount.Code).NotTo(BeNil())
			})
		})
	})
})
