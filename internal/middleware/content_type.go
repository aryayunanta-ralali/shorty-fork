// Package middleware
package middleware

import (
	"net/http"
	"strings"
	"fmt"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
)

// ValidateContentType header
func ValidateContentType(r *http.Request, conf *appctx.Config) string {

	if ct := strings.ToLower(r.Header.Get(`Content-Type`)) ; ct != `application/json` {
		logger.Warn(fmt.Sprintf("[middleware] invalid content-type %s", ct ))

		return consts.ResponseUnprocessableEntity
	}


	return consts.MiddlewarePassed
}
