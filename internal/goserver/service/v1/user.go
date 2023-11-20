package v1

import (
	v1 "baseframe/api/goserver/v1"
	"baseframe/internal/pkg/code"
	"baseframe/internal/pkg/models"
	"baseframe/pkg/errors"
	"context"
	"github.com/jinzhu/copier"
)

type UserRepo interface {
	Create(ctx context.Context, u *models.UserModel) error
	GetByUserName(ctx context.Context, username string) (*models.UserModel, error)
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

func (s *UserService) Login(ctx context.Context, r *v1.LoginReq) (*models.UserModel, error) {
	user, err := s.repo.GetByUserName(ctx, r.Username)
	if err != nil {
		return nil, errors.WithCode(code.ErrUserNotFound, err.Error())
	}

	if err := user.Compare(r.Password); err != nil {
		return nil, err
	}

	return user, nil
}
