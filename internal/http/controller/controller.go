package controller

import (
	"github.com/foxfurry/go_aggregator/internal/domain/dto"
	"github.com/foxfurry/go_aggregator/internal/service/supervisor"
	"github.com/gin-gonic/gin"
)

type IController interface {
	RegisterDeliveryRoutes(c *gin.Engine)
}

type aggregatorController struct {
	supervisor supervisor.ISupervisor
}

func NewAggregatorController() IController {
	return &aggregatorController{
		supervisor: supervisor.NewAggregatorSupervisor(),
	}
}

func (ctrl *aggregatorController) RegisterDeliveryRoutes(c *gin.Engine) {
	c.GET("/menu", ctrl.getAllMenus)
	c.POST("/register", ctrl.register)
	c.POST("/order", ctrl.order)
}

func (ctrl *aggregatorController) getAllMenus(c *gin.Context) {
	menus, err := ctrl.supervisor.GetMenus()
	if err != nil {
		c.JSON(500, gin.H{"Internal error": err})
		return
	}

	c.JSON(200, menus)
}

func (ctrl *aggregatorController) register(c *gin.Context) {
	var currentRegister dto.RestaurantRegister

	if err := c.ShouldBindJSON(&currentRegister); err != nil {
		c.JSON(500, gin.H{"Internal error": err})
		return
	}

	err := ctrl.supervisor.RegisterRestaurant(currentRegister)
	if err != nil {
		c.JSON(500, gin.H{"Internal error": err.Error()})
		return
	}

	c.JSON(200, "Successfully registered")
}

func (ctrl *aggregatorController) order(c *gin.Context) {
	order := new(dto.Order)

	if err := c.ShouldBindJSON(order); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	if err := ctrl.supervisor.Order(order); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, "Successfully made order")
}
