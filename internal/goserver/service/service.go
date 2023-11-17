package service

import (
	"baseframe/internal/goserver/service/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(v1.NewUserService)
