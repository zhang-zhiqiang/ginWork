package user

import (
	v1 "baseframe/api/goserver/v1"
	"baseframe/pkg/log"
	"context"
	"errors"
	"github.com/jinzhu/copier"
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

type UserRepo interface {
	Create(ctx context.Context, u *UserModel) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, r *v1.CreateUserReq) error {
	um := &UserModel{}
	if err := copier.Copy(um, r); err != nil {
		log.Warnf("copier 复制值失败 %v", err)
		return errors.New("创建失败")
	}
	return s.repo.Create(ctx, um)
}
