package entities

import (
	"time"

	"github.com/google/uuid"
)

type Discount struct {
	ID           uuid.UUID
	RestaurantID uuid.UUID
	Percentage int
	StartDate time.Time
	EndDate time.Time
	Description string
}