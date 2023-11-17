package data

import (
	"baseframe/internal/goserver/service/v1"
	"baseframe/internal/pkg/models"
	"context"
	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

func NewUserData(db *gorm.DB) v1.UserRepo {
	return &UserData{db: db}
}

func (ud *UserData) Create(ctx context.Context, user *models.UserModel) error {
	return ud.db.Create(&user).Error
}
