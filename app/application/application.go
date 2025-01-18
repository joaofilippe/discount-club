package application

import "github.com/joaofilippe/discount-club/app/domain/iservices"

type Application struct {
	discountService iservices.IDiscount
}

func (a *Application) DiscountService() iservices.IDiscount {
	return a.discountService
}
