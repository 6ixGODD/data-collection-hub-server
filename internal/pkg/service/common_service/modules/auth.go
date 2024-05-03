package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type AuthService interface {
}

type AuthServiceImpl struct {
	service     *service.Core
	userDao     dao.UserDao
	loginLogDao dao.LoginLogDao
}

func NewAuthService(s *service.Core, userDao dao.UserDao, loginLogDao dao.LoginLogDao) AuthService {
	return &AuthServiceImpl{
		service:     s,
		userDao:     userDao,
		loginLogDao: loginLogDao,
	}
}
