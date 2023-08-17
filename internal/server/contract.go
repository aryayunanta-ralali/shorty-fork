// Package router
package server

import (
	"context"
	"net/http"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	ucase "github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
)

// httpHandlerFunc abstraction for http handler
type httpHandlerFunc func(request *http.Request, svc ucase.UseCase, conf *appctx.Config) appctx.Response

// Server contract
type Server interface {
	Run(context.Context) error
	Done()
	Config() *appctx.Config
}
