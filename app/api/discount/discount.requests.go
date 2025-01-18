package discountwebserver

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreateRequest struct {
	RestaurantIDStr string `json:"restaurant_id"`
	RestaurantID    uuid.UUID
	UserIDStr       string `json:"user_id"`
	UserID          uuid.UUID
	Percentage      int    `json:"percentage"`
	StartDateStr    string `json:"start_date"`
	StartDate       time.Time
	EndDateStr      string `json:"end_date"`
	EndDate         time.Time 
	Times           int    `json:"times"`
	Description     string `json:"description"`
}

func (cr *CreateRequest) parseDates() error {
	startDate, err := time.Parse(time.DateOnly, cr.StartDateStr)
	if err != nil {
		return fmt.Errorf("invalid start date: %w", err)
	}

	endDate, err := time.Parse(time.DateOnly, cr.EndDateStr)
	if err != nil {
		return fmt.Errorf("invalid end date: %w", err)
	}

	cr.StartDate = startDate
	cr.EndDate = endDate

	return nil
}

func (cr *CreateRequest) parseUUIDs() error {
	restaurantID, err := uuid.Parse(cr.RestaurantIDStr)
	if err != nil {
		return fmt.Errorf("invalid restaurant id: %w", err)
	}

	userID, err := uuid.Parse(cr.UserIDStr)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}

	cr.RestaurantID = restaurantID
	cr.UserID = userID

	return nil
}


func (cr *CreateRequest) Parse() (*CreateRequest, error) {
	newRequest := &CreateRequest{
		RestaurantIDStr: cr.RestaurantIDStr,
		RestaurantID:    cr.RestaurantID,
		UserIDStr:       cr.UserIDStr,
		UserID:          cr.UserID,
		Percentage:      cr.Percentage,
		StartDateStr:    cr.StartDateStr,
		StartDate:       cr.StartDate,
		EndDateStr:      cr.EndDateStr,
		EndDate:         cr.EndDate,
		Times:           cr.Times,
		Description:     cr.Description,
	}

	err := newRequest.parseDates()
	if err != nil {
		return &CreateRequest{}, err
	}

	err = newRequest.parseUUIDs()
	if err != nil {
		return &CreateRequest{}, err
	}

	return newRequest, nil
}
