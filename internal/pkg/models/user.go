package models

import (
	"baseframe/pkg/log"
	"baseframe/pkg/utils"
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
	Status    uint8          `gorm:"column:status"`
	CreatedAt time.Time      `gorm:"column:createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt"`
}

const (
	UserStatusNormal uint8 = iota + 1
	UserStatusDelete
)

var UserStatusDescriptions = map[uint8]string{
	UserStatusNormal: "正常",
	UserStatusDelete: "退出",
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) BeforeCreate(_ *gorm.DB) error {
	if u.Status == 0 {
		u.Status = UserStatusNormal
	}
	hashedPwd, err := utils.Encrypt(u.Password)
	if err != nil {
		log.Panicf("密码加密有误 %v", err)
		return err
	}
	u.Password = hashedPwd
	return nil
}
