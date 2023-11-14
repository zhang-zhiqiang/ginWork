package user

import "baseframe/internal/goserver/service/user"

type UserController struct {
	us *user.UserService
}

func NewUserController(us *user.UserService) *UserController {
	return &UserController{us}
}
