package common

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type AuthService interface {
}

type AuthServiceImpl struct {
	Service     *service.Service
	UserDao     dao.UserDao
	LoginLogDao dao.LoginLogDao
}

func NewAuthService(s *service.Service, userDaoImpl *dao.UserDaoImpl, loginLogDaoImpl *dao.LoginLogDaoImpl) AuthService {
	return &AuthServiceImpl{
		Service:     s,
		UserDao:     userDaoImpl,
		LoginLogDao: loginLogDaoImpl,
	}
}
