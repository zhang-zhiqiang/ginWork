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

func (ud *UserData) GetByUserName(ctx context.Context, username string) (*models.UserModel, error) {
	user := &models.UserModel{}
	err := ud.db.Where("username = ? and status = ?", username, models.UserStatusNormal).First(&user).Error
	return user, err
}
