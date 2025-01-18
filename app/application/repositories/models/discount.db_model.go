package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/discount-club/app/domain/entities"
)

type Discount struct {
	ID           uuid.UUID `db:"id"`
	RestaurantID uuid.UUID `db:"restaurant_id"`
	UserID       uuid.UUID `db:"user_id"`
	Code         string    `db:"code"`
	Percentage   int       `db:"percentage"`
	StartDate    time.Time `db:"start_date"`
	EndDate      time.Time `db:"end_date"`
	Times        int       `db:"times"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	DeletedAt    time.Time `db:"deleted_at"`
	Active       bool      `db:"active"`
}

func (d *Discount) FromEntity(entity *entities.Discount) {
	d.ID = entity.ID()
	d.RestaurantID = entity.RestaurantID()
	d.UserID = entity.UserID()
	d.Code = entity.Code()
	d.Percentage = entity.Percentage()
	d.StartDate = entity.StartDate()
	d.EndDate = entity.EndDate()
	d.Times = entity.Times()
	d.Description = entity.Description()
	d.CreatedAt = entity.CreatedAt()
	d.UpdatedAt = entity.UpdatedAt()
	d.DeletedAt = entity.DeletedAt()
	d.Active = entity.Active()
}
