package responses

import (
	"github.com/joaofilippe/discount-club/app/domain/entities"
)

type UserDTO struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	UserType  string `json:"user_type"`
	CreatedAt string `json:"created_at"`
	Active    bool   `json:"active"`
}

func NewUserDTOFromEntity(entity entities.User) UserDTO {
	return UserDTO{
		ID:        entity.ID().String(),
		Email:     entity.Email(),
		UserType:  entity.UserType().String(),
		CreatedAt: entity.CreatedAt().Format("2006-01-02 15:04:05"),
		Active:    entity.Active(),
	}
}
