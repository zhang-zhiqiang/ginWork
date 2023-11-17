package v1

import (
	v1 "baseframe/api/goserver/v1"
	"baseframe/internal/pkg/models"
	"context"
	"errors"
	"github.com/jinzhu/copier"
)

type UserRepo interface {
	Create(ctx context.Context, u *models.UserModel) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, r *v1.CreateUserReq) error {
	um := &models.UserModel{}
	if err := copier.Copy(um, r); err != nil {
		return errors.New("创建失败")
	}
	return s.repo.Create(ctx, um)
}
