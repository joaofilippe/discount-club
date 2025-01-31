package irepositories

import "github.com/joaofilippe/discount-club/app/domain/entities"

type User interface {
	Save(user *entities.User) error
}
