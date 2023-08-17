// Package example
package example

import (
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"

	"github.com/aryayunanta-ralali/shorty/pkg/logger"
)

type exampleList struct {
	repo repositories.Example
}

func NewExampleList(repo repositories.Example) contract.UseCase {
	return &exampleList{repo: repo}
}

// Serve partner list data
func (u *exampleList) Serve(data *appctx.Data) appctx.Response {

	p, e := u.repo.Find(data.Request.Context())

	if e != nil {
		logger.Error(logger.SetMessageFormat("[example-list] %v", e.Error()))

		return appctx.Response{
			Name: consts.ResponseInternalFailure,
		}
	}

	return appctx.Response{
		Name:    consts.ResponseSuccess,
		Message: "ok",
		Data:    p,
	}
}
