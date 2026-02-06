package irepositories

import (
	"github.com/joaofilippe/discount-club/app/domain/entities"
)

type IUser interface {
	Save(user *entities.User) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
}
