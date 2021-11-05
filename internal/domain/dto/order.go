package dto

type Order struct {
	ClientID int `json:"client_id"`
	RestaurantID int `json:"restaurant_id"`
	Items []int `json:"items"`
	Priority int `json:"priority"`
	MaxWait int `json:"max_wait"`
}
