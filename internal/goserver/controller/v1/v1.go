package v1

import (
	"baseframe/internal/goserver/controller/v1/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(user.NewUserController)
