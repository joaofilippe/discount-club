package application

import "github.com/joaofilippe/discount-club/app/domain/iservices"

type IApplication interface {
	DiscountService() iservices.IDiscount
	UserService() iservices.IUser
}

type Application struct {
	discountService iservices.IDiscount
	userService     iservices.IUser
}

func New(discountService iservices.IDiscount, userService iservices.IUser) *Application {
	return &Application{
		discountService: discountService,
		userService:     userService,
	}
}

func (a *Application) DiscountService() iservices.IDiscount {
	return a.discountService
}

func (a *Application) UserService() iservices.IUser {
	return a.userService
}
