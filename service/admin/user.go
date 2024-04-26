package admin

import (
	dao "data-collection-hub-server/dal/modules"
	"data-collection-hub-server/service"
)

type UserService interface {
}

type UserServiceImpl struct {
	*service.Service
	dao.UserDao
}

func NewUserService(s *service.Service, userDaoImpl *dao.UserDaoImpl) UserService {
	return &UserServiceImpl{
		Service: s,
		UserDao: userDaoImpl,
	}
}
