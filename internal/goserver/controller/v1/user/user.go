package user

import (
	"baseframe/internal/goserver/service/v1"
)

type UserController struct {
	us *v1.UserService
}

func NewUserController(us *v1.UserService) *UserController {
	return &UserController{us}
}
