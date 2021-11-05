package main

import (
	"context"
	"github.com/foxfurry/go_aggregator/application"
	"github.com/foxfurry/go_aggregator/internal/infrastracture/config"
	"github.com/foxfurry/go_aggregator/internal/infrastracture/profiler"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.LoadConfig()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	app := application.Create(ctx)

	go app.Start()
	go profiler.Start(ctx)

	<-sigs
	cancel()
	app.Shutdown()
}
