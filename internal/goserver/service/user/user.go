package user

import (
	v1 "baseframe/api/goserver/v1"
	"context"
	log "github.com/golang/glog"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        int64          `gorm:"column:id;primary_key" json:"id"`
	Username  string         `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password  string         `json:"password,omitempty" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Nickname  string         `json:"nickname" gorm:"column:nickname" binding:"required" validate:"required,min=1,max=30"`
	Email     string         `json:"email" gorm:"column:email" binding:"required" validate:"required,email,min=1,max=100"`
	Phone     string         `json:"phone" gorm:"column:phone" binding:"required" validate:"required,phone,min=1,max=16"`
	CreatedAt time.Time      `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt" json:"deletedAt"`
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
	log.Info("user service")
	um := &UserModel{}
	s.repo.Create(ctx, um)
	return nil
}
