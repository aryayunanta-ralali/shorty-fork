// Package router
package router

import (
	"net/http"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/routerkit"
)

// httpHandlerFunc is a contract http handler for router
type httpHandlerFunc func(request *http.Request, svc contract.UseCase, conf *appctx.Config) appctx.Response

// Router is a contract router and must implement this interface
type Router interface {
	Route() *routerkit.Router
}
