package application

import "github.com/joaofilippe/discount-club/app/domain/iservices"

type IApplication interface {
	DiscountService() iservices.IDiscount
}

type Application struct {
	discountService iservices.IDiscount
}

func New(discountService iservices.IDiscount) *Application {
	return &Application{
		discountService: discountService,
	}
}

func (a *Application) DiscountService() iservices.IDiscount {
	return a.discountService
}
