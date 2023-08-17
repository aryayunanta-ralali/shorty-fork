// Package example
package example

import (
	"github.com/thedevsaddam/govalidator"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/internal/entity"
	"github.com/aryayunanta-ralali/shorty/internal/presentations"
	"github.com/aryayunanta-ralali/shorty/internal/repositories"
	"github.com/aryayunanta-ralali/shorty/internal/ucase/contract"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/util"
)

type exampleCreate struct {
	repo repositories.Example
}

// NewPartnerCreate initialize partner cerator
func NewPartnerCreate(repo repositories.Example) contract.UseCase {
	return &exampleCreate{repo: repo}
}

// Serve partner list data
func (u *exampleCreate) Serve(data *appctx.Data) appctx.Response {

	req := presentations.ExampleCreate{}

	e := data.Cast(&req)

	if e != nil {
		logger.Error(logger.SetMessageFormat("[example-create] parsing body request error: %s", e.Error()))
		return appctx.Response{
			Name: consts.ResponseValidationFailure,
		}
	}

	fl := []logger.Field{
		logger.Any("request", req),
	}

	rules := govalidator.MapData{
		"name":    []string{"required", "between:3,50"},
		"email":   []string{"required", "email"},
		"address": []string{"required"},
		"phone":   []string{"phone_number"},
	}

	opts := govalidator.Options{
		Data:  &req,  // request object
		Rules: rules, // rules map
	}

	v := govalidator.New(opts)
	ev := v.ValidateStruct()

	if len(ev) != 0 {
		logger.Warn(
			logger.SetMessageFormat("[example-create] validate request param err: %s", util.DumpToString(ev)),
			fl...)

		return appctx.Response{
			Name:   consts.ResponseValidationFailure,
			Errors: ev,
		}
	}

	_, e = u.repo.Upsert(data.Request.Context(), entity.Example{
		Name:    req.Name,
		Address: req.Address,
		Email:   req.Email,
		Phone:   req.Phone,
	})

	if e != nil {
		logger.Error(logger.SetMessageFormat("[example-create] %v", e.Error()))

		return appctx.Response{
			Name: consts.ResponseInternalFailure,
		}
	}

	return appctx.Response{
		Name:    consts.ResponseSuccess,
		Message: "ok",
	}
}
