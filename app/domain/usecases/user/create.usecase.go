package userusecases

import "github.com/joaofilippe/discount-club/app/domain/irepositories"

type CreateUserUseCase struct {
	userRepo irepositories.User
}

func NewCreateUserUseCase(userRepo irepositories.User) *CreateUserUseCase {
	if userRepo == nil {
		panic("userRepo is required")
	}

	return &CreateUserUseCase{
		userRepo: userRepo,
	}
}