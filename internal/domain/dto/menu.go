package dto

type Menu struct {
	Restaurants int `json:"restaurants"`
	RestaurantsData []Restaurant `json:"restaurants_data"`
}
