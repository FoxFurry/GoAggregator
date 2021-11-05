package supervisor

import (
	"errors"
	"github.com/foxfurry/go_aggregator/internal/domain/dto"
	"github.com/foxfurry/go_aggregator/internal/domain/entity"
	"github.com/foxfurry/go_aggregator/internal/infrastracture/logger"
	"github.com/foxfurry/go_aggregator/internal/service/supervisor/gateway"
)

type ISupervisor interface {
	RegisterRestaurant(register dto.RestaurantRegister) error
	GetMenus() (*dto.Menu, error)
	Order(order *dto.Order) error
}

type aggregatorSupervisor struct {
	restaurants map[int]entity.Restaurant
	gateway gateway.IGateway
}

func NewAggregatorSupervisor() ISupervisor{
	return &aggregatorSupervisor{
		restaurants: make(map[int]entity.Restaurant),
		gateway: gateway.NewAggregatorGateway(),
	}
}

func (s *aggregatorSupervisor) RegisterRestaurant(register dto.RestaurantRegister) error {
	if _, ok := s.restaurants[register.RestaurantID]; ok {	// If restaurant already exists
		logger.LogSuperF("Restaurant %s tried to allocate ID %d, which is already allocated by restaurant %s", register.Name, register.RestaurantID, s.restaurants[register.RestaurantID].Name)
		return errors.New("restaurant with such ID already exists")
	}else{
		s.restaurants[register.RestaurantID] = register.Restaurant
	}

	logger.LogSuperF("Registered restaurant %s with ID %d", register.Name, register.RestaurantID)
	return nil
}

func (s *aggregatorSupervisor) GetMenus() (*dto.Menu, error) {
	if len(s.restaurants) == 0 {
		return nil, nil
	}

	result := new(dto.Menu)
	result.Restaurants = len(s.restaurants)
	for _, val := range s.restaurants {
		restDto := dto.Restaurant{
			Name:           val.Name,
			MenuItems:      val.MenuItems,
			RestaurantMenu: val.Menu,
		}
		result.RestaurantsData = append(result.RestaurantsData, restDto)
	}


	return result, nil
}

func (s *aggregatorSupervisor) Order(order *dto.Order) error{
	restaurantHost := s.restaurants[order.RestaurantID].Address
	logger.LogSuperF("Received order from client %d for restaurant %s [%s]", order.ClientID, s.restaurants[order.RestaurantID].Name,restaurantHost)

	s.gateway.SendOrder(order, restaurantHost)

	return nil
}

