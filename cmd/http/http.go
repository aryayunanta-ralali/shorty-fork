package http

import (
	"context"

	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/server"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
)

// Start function handler starting http listener
func Start(ctx context.Context) {

	serve := server.NewHTTPServer()
	defer serve.Done()
	logger.Info(logger.SetMessageFormat("starting shorty services... %d", serve.Config().App.Port), logger.EventName(consts.LogEventNameServiceStarting))

	if err := serve.Run(ctx); err != nil {
		logger.Warn(logger.SetMessageFormat("service stopped, err:%s", err.Error()), logger.EventName(consts.LogEventNameServiceStarting))
	}

	return
}
