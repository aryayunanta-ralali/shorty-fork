// Package bootstrap
package bootstrap

import (
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

func RegistryLogger(cfg *appctx.Config) {

	lc := logger.Config{
		URL:         cfg.Logger.URL,
		Environment: util.EnvironmentTransform(cfg.App.Env),
		Debug:       cfg.App.Debug,
		Level:       cfg.Logger.Level,
	}

	h, e := logger.NewSentryHook(lc)

	if e != nil {
		logger.Fatal("log sentry failed to initialize Sentry")
	}

	logger.Setup(lc, h)
}
