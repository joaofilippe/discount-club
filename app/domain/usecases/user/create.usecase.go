package userusecases

import (
	usererrors "github.com/joaofilippe/discount-club/app/common/myerrors/user"
	"github.com/joaofilippe/discount-club/app/domain/irepositories"
)

type CreateUserUseCase struct {
	userRepo irepositories.User
}

func NewCreateUserUseCase(userRepo irepositories.User) (*CreateUserUseCase, error) {
	if userRepo == nil {
		return &CreateUserUseCase{}, usererrors.ErrNilUserRepo
	}

	return &CreateUserUseCase{
		userRepo: userRepo,
	}, nil
}

func (uc *CreateUserUseCase) Execute() error {
	return nil
}