package data

import (
	"baseframe/internal/goserver/service/user"
	"baseframe/pkg/log"
	"context"
)

type UserData struct {
	data *Data
}

func NewUserData(data *Data) user.UserRepo {
	return &UserData{data: data}
}

func (ud *UserData) Create(ctx context.Context, u *user.UserModel) error {
	log.Info("user data")
	return nil
}
