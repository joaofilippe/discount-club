package entities

import (
	"time"

	"github.com/google/uuid"
)

type Discount struct {
	id           uuid.UUID
	restaurantID uuid.UUID
	userID       uuid.UUID
	code         string
	percentage   int
	startDate    time.Time
	endDate      time.Time
	times        int
	description  string
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    time.Time
	active       bool
}

func NewDiscount(
	restaurantID, userID uuid.UUID,
	percentage int,
	startDate, endDate time.Time,
	times int,
	description string,
) *Discount {
	return &Discount{
		id:           uuid.UUID{},
		restaurantID: restaurantID,
		userID:       userID,
		code:         "",
		percentage:   percentage,
		startDate:    startDate,
		endDate:      endDate,
		times:        times,
		description:  description,
		active:       true,
	}
}

func (d *Discount) ID() uuid.UUID {
	return d.id
}

func (d *Discount) RestaurantID() uuid.UUID {
	return d.restaurantID
}

func (d *Discount) UserID() uuid.UUID {
	return d.userID
}

func (d *Discount) Code() string {
	return d.code
}

func (d *Discount) Percentage() int {
	return d.percentage
}

func (d *Discount) StartDate() time.Time {
	return d.startDate
}

func (d *Discount) EndDate() time.Time {
	return d.endDate
}

func (d *Discount) Times() int {
	return d.times
}

func (d *Discount) Description() string {
	return d.description
}

func (d *Discount) CreatedAt() time.Time {
	return d.createdAt
}

func (d *Discount) UpdatedAt() time.Time {
	return d.updatedAt
}

func (d *Discount) DeletedAt() time.Time {
	return d.deletedAt
}

func (d *Discount) Active() bool {
	return d.active
}

func (d *Discount) SetID(id uuid.UUID) {
	d.id = id
}

func (d *Discount) SetCode(code string) {
	d.code = code
}
