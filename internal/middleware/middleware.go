// Package middleware
package middleware

import (
	"net/http"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
)

// MiddlewareFunc is contract for middleware and must implement this type for http if need middleware http request
type MiddlewareFunc func(r *http.Request, conf *appctx.Config) string

// FilterFunc is a iterator resolver in each middleware registered
func FilterFunc(conf *appctx.Config, r *http.Request, mfs []MiddlewareFunc) string {
	for _, mf := range mfs {
		if status := mf(r, conf); status != consts.MiddlewarePassed {
			return status
		}
	}

	return consts.MiddlewarePassed
}
