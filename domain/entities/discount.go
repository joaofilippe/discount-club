package entities

import (
	"time"

	"github.com/google/uuid"
)

type Discount struct {
	ID           uuid.UUID
	RestaurantID uuid.UUID
	UserID       uuid.UUID
	Code         string
	Percentage   int
	StartDate    time.Time
	EndDate      time.Time
	Times        int
	Description  string
}
