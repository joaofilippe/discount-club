package discountusecases_test

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	discounterrors "github.com/joaofilippe/discount-club/app/common/myerrors/discount"
	"github.com/joaofilippe/discount-club/app/domain/entities"
	discountusecases "github.com/joaofilippe/discount-club/app/domain/usecases/discount"
	"github.com/joaofilippe/discount-club/mocks"
)

var _ = Describe("Create.Usecase", func() {
	var (
		discountRepoMock *mocks.Discount
		discountUseCase  *discountusecases.CreateUseCase
		request          *entities.Discount
		err              error
	)

	Describe("NewCreateUseCase", func() {
		Context("whith a valid discount repository", func() {
			BeforeEach(func() {
				discountRepoMock = &mocks.Discount{}
				discountUseCase, err = discountusecases.NewCreateUseCase(discountRepoMock)
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
				Expect(err).To(Equal(discounterrors.ErrNilDiscountRepo))
			})
		})
	})

	Describe("Execute", func() {
		BeforeEach(func() {
			discountRepoMock = &mocks.Discount{}

			discountUseCase, _ = discountusecases.NewCreateUseCase(discountRepoMock)

			startTime, _ := time.Parse("2006-01-02", "2021-01-31")
			endTime := startTime.Add(time.Hour * 24 * 31)
			request, _ = entities.NewDiscount(
				uuid.New(),
				uuid.New(),
				10,
				startTime,
				endTime,
				1,
				"test discount",
			)

			discountRepoMock.On("Save", mock.Anything).Return(nil)
		})

		Context("whith a valid discount", func() {
			BeforeEach(func(){

			})

			It("should return no error", func() {
				request, err = discountUseCase.Execute(request)
				if err != nil {
					fmt.Println(err)
				}

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(request.Code).NotTo(BeNil())
			})
		})

		Context("with a nil request", func() {
			BeforeEach(func() {
				request = nil
			})

			_, err := discountUseCase.Execute(request)
			It("should return ErrNoDiscountProvided", func() {
				Expect(err).NotTo(BeNil())
				Expect(err).To(Equal(discounterrors.ErrNoDiscountProvided))
			})
		})
	})
})
