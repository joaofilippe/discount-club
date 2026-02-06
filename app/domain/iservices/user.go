package iservices

import "github.com/joaofilippe/discount-club/app/domain/entities"

type IUser interface {
	Create(user *entities.User) (*entities.User, error)
	Login(email, password string) (string, error)
}
