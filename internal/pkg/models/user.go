package models

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        int64          `gorm:"column:id;primary_key"`
	Username  string         `gorm:"column:username;not null"`
	Password  string         `gorm:"column:password;not null"`
	Nickname  string         `gorm:"column:nickname"`
	Email     string         `gorm:"column:email"`
	Phone     string         `gorm:"column:phone"`
	CreatedAt time.Time      `gorm:"column:createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt"`
}

func (u *UserModel) TableName() string {
	return "user"
}
