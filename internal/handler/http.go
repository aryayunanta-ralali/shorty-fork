// Package handler
package handler

import (
	"net/http"
	"context"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/msg"
)


// HttpRequest handler func wrapper
func HttpRequest(request *http.Request, svc contract.UseCase, conf *appctx.Config) appctx.Response {
	if !msg.GetAvailableLang(consts.ResponseSuccess, request.Header.Get(consts.HeaderLanguageKey)) {
		request.Header.Set(consts.HeaderLanguageKey, conf.App.DefaultLang)
	}

	ctx := context.WithValue(request.Context(), consts.CtxLang, request.Header.Get(consts.HeaderLanguageKey))

	req := request.WithContext(ctx)

	data := &appctx.Data{
		Request:     req,
		Config:      conf,
		ServiceType: consts.ServiceTypeHTTP,
	}

	return svc.Serve(data)
}
