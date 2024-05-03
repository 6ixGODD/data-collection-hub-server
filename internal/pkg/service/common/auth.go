package common

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type AuthService interface {
}

type authServiceImpl struct {
	service     *service.Service
	userDao     dao.UserDao
	loginLogDao dao.LoginLogDao
}

func NewAuthService(s *service.Service, userDao dao.UserDao, loginLogDao dao.LoginLogDao) AuthService {
	return &authServiceImpl{
		service:     s,
		userDao:     userDao,
		loginLogDao: loginLogDao,
	}
}
