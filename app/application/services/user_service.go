package services

import (
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
)

type UserService struct {
	userRepo irepositories.IUser
}

func NewUserService(userRepo irepositories.IUser) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(user *entities.User) (*entities.User, error) {
	// Here we could add business logic, e.g., hashing password, checking duplicates, etc.
	// For now, simple pass-through.
	return s.userRepo.Save(user)
}
