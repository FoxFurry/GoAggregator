package application

import (
	"context"
	"github.com/foxfurry/go_aggregator/internal/http/controller"
	"github.com/foxfurry/go_aggregator/internal/infrastracture/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type IApp interface {
	Start()
	Shutdown()
}

type aggregatorApp struct {
	server *http.Server
}

func Create(ctx context.Context) IApp {
	appHandler := gin.New()

	ctrl := controller.NewAggregatorController()
	ctrl.RegisterDeliveryRoutes(appHandler)

	app := aggregatorApp{
		server: &http.Server{
			Addr:              viper.GetString("aggregator_host"),
			Handler:           appHandler,
		},
	}

	return &app
}

func (d *aggregatorApp) Start() {
	logger.LogMessage("Starting aggregator server")

	if err := d.server.ListenAndServe(); err != http.ErrServerClosed {
		logger.LogPanicF("Unexpected error while running server: %v", err)
	}
}

func (d *aggregatorApp) Shutdown() {
	if err := d.server.Shutdown(context.Background()); err != nil {
		logger.LogPanicF("Unexpected error while closing server: %v", err)
	}
	logger.LogMessage("Server terminated successfully")
}
