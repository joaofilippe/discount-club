package discountwebserver

import "github.com/joaofilippe/discount-club/app/domain/entities"

type DiscountDTO struct {
	ID           string `json:"id"`
	RestaurantID string `json:"restaurant_id"`
	UserID       string `json:"user_id"`
	Percentage   int    `json:"percentage"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
	Times        int    `json:"times"`
	Description  string `json:"description"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
	Active       bool   `json:"active"`
}

func NewDiscountDTOFromEntity(entity entities.Discount) DiscountDTO {
	return DiscountDTO{
		ID:           entity.ID().String(),
		RestaurantID: entity.RestaurantID().String(),
		UserID:       entity.UserID().String(),
		Percentage:   entity.Percentage(),
		StartDate:    entity.StartDate().Format("2006-01-02"),
		EndDate:      entity.EndDate().Format("2006-01-02"),
		Times:        entity.Times(),
		Description:  entity.Description(),
		CreatedAt:    entity.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:    entity.UpdatedAt().Format("2006-01-02 15:04:05"),
		DeletedAt:    entity.DeletedAt().Format("2006-01-02 15:04:05"),
		Active:       entity.Active(),
	}
}
