package dto

import "github.com/foxfurry/go_aggregator/internal/domain/entity"

type RestaurantRegister struct {
	RestaurantID int `json:"restaurant_id"`
	entity.Restaurant
}
