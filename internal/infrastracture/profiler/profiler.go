package profiler

import (
	"context"
	"github.com/foxfurry/go_aggregator/internal/infrastracture/logger"
	"net/http"
	_ "net/http/pprof"
)

func Start(ctx context.Context) {
	logger.LogMessage("Starting aggregator profiler. Access http://localhost:6060/debug/pprof/ for more useful data!")
	go logger.LogError(http.ListenAndServe("localhost:6060", nil).Error())
	<-ctx.Done()
	logger.LogMessage("Shutting down profiler server!")
	return
}