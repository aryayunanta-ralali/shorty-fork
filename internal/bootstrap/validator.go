// Package bootstrap
package bootstrap

import (
	"github.com/thedevsaddam/govalidator"

	"github.com/aryayunanta-ralali/shorty/internal/appctx"
	"github.com/aryayunanta-ralali/shorty/internal/validator"
)

// RegistryValidatorRules initialize validation rules
func RegistryValidatorRules(cfg *appctx.Config) {
	for k, f := range validator.Rules(cfg.App.Env) {
		govalidator.AddCustomRule(k, f)
	}
}
