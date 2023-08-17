// Package bootstrap
package bootstrap

import (
	"github.com/aryayunanta-ralali/shorty/internal/consts"
	"github.com/aryayunanta-ralali/shorty/pkg/logger"
	"github.com/aryayunanta-ralali/shorty/pkg/msg"
)

func RegistryMessage()  {
	err := msg.Setup("msg.yaml", consts.ConfigPath)
	if err != nil {
		logger.Fatal(logger.SetMessageFormat("file message multi language load error %s", err.Error()))
	}

}
