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
) (*Discount, error) {
	return &Discount{
		id:           uuid.New(),
		restaurantID: restaurantID,
		userID:       userID,
		code:         "",
		percentage:   percentage,
		startDate:    startDate,
		endDate:      endDate,
		times:        times,
		description:  description,
		active:       true,
	}, nil
}

func (d *Discount) CopyWith(
	restaurantID, userID *uuid.UUID,
	code *string,
	percentage *int,
	startDate, endDate *time.Time,
	times *int,
	description *string,
	createdAt *time.Time,
	updatedAt *time.Time,
	deletedAt *time.Time,
	active *bool,
) *Discount {
	newDiscount := &Discount{
		id:           d.id,
		restaurantID: d.restaurantID,
		userID:       d.userID,
		code:         d.code,
		percentage:   d.percentage,
		startDate:    d.startDate,
		endDate:      d.endDate,
		times:        d.times,
		description:  d.description,
		createdAt:    d.createdAt,
		updatedAt:    d.updatedAt,
		deletedAt:    d.deletedAt,
		active:       d.active,
	}

	if restaurantID != nil {
		newDiscount.restaurantID = *restaurantID
	}

	if userID != nil {
		newDiscount.userID = *userID
	}

	if code != nil {
		newDiscount.code = *code
	}

	if percentage != nil {
		newDiscount.percentage = *percentage
	}

	if startDate != nil {
		newDiscount.startDate = *startDate
	}

	if endDate != nil {
		newDiscount.endDate = *endDate
	}

	if times != nil {
		newDiscount.times = *times
	}

	if description != nil {
		newDiscount.description = *description
	}

	if createdAt != nil {
		newDiscount.createdAt = *createdAt
	}

	if updatedAt != nil {
		newDiscount.updatedAt = *updatedAt
	}

	if deletedAt != nil {
		newDiscount.deletedAt = *deletedAt
	}

	if active != nil {
		newDiscount.active = *active
	}

	return newDiscount
}

func (d *Discount) CopyWithNewCode(code *string) *Discount {
	return d.CopyWith(
		nil,
		nil,
		code,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
}

func NewEmptyDiscount() *Discount {
	return &Discount{}
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

func (d *Discount) IsEmpty() bool {
	return d.id == uuid.UUID{} &&
		d.code == "" &&
		d.percentage == 0 &&
		d.startDate == time.Time{} &&
		d.endDate == time.Time{} &&
		d.times == 0 &&
		d.description == "" &&
		d.createdAt == time.Time{} &&
		d.updatedAt == time.Time{} &&
		d.deletedAt == time.Time{} &&
		!d.active
}
