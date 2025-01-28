package irepositories

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/discount-club/app/domain/entities"
)

type Discount interface {
	Save(discount *entities.Discount) error
	GetByID(id uuid.UUID) (*entities.Discount, error)
	GetByCode(code string) (*entities.Discount, error)
}
