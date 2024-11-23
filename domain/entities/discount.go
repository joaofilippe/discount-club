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
	}
}
