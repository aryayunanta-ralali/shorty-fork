// Package example
package example

import (
	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	
	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"

	"github.com/aryayunanta-ralali/shorty/pkg/logger"
)

type exampleDelete struct {
	repo repositories.Example
}

func NewExampleDelete(repo repositories.Example) contract.UseCase {
	return &exampleDelete{repo: repo}
}

// Serve partner list data
func (u *exampleDelete) Serve(data *appctx.Data) appctx.Response {

	id := mux.Vars(data.Request)["id"]

	 e := u.repo.Delete(data.Request.Context(), cast.ToUint64(id))

	if e != nil {
		logger.Error(logger.SetMessageFormat("[example-delete] %v", e.Error()))

		return appctx.Response{
			Name: consts.ResponseInternalFailure,
		}
	}

	return appctx.Response{
		Name:    consts.ResponseSuccess,
	}
}
