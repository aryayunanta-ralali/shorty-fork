package ucase

import (
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
)

type healthCheck struct {
}

func NewHealthCheck() contract.UseCase {
	return &healthCheck{}
}

func (u *healthCheck) Serve(*appctx.Data) appctx.Response {
	return appctx.Response{
		Name:    consts.ResponseSuccess,
		Message: "ok",
	}
}
