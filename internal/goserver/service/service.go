package service

import (
	"baseframe/internal/goserver/service/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(user.NewUserService)
