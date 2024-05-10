package mods

import (
	"context"

	dao "data-collection-hub-server/internal/pkg/dal/mods"
	"data-collection-hub-server/internal/pkg/schema/common"
	"data-collection-hub-server/internal/pkg/service"
)

type AuthService interface {
	Login(ctx context.Context, email, password *string) (*common.LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken *string) (*common.RefreshTokenResponse, error)
	Logout(ctx context.Context) error
	ChangePassword(ctx context.Context, oldPassword, newPassword *string) error
}

type AuthServiceImpl struct {
	service     *service.Service
	userDao     dao.UserDao
	loginLogDao dao.LoginLogDao
}

func (a AuthServiceImpl) Login(ctx context.Context, email, password *string) (*common.LoginResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) RefreshToken(ctx context.Context, refreshToken *string) (*common.RefreshTokenResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) Logout(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (a AuthServiceImpl) ChangePassword(ctx context.Context, oldPassword, newPassword *string) error {
	// TODO implement me
	panic("implement me")
}

func NewAuthService(s *service.Service, userDao dao.UserDao, loginLogDao dao.LoginLogDao) AuthService {
	return &AuthServiceImpl{
		service:     s,
		userDao:     userDao,
		loginLogDao: loginLogDao,
	}
}
