package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type UserService interface {
}

type UserServiceImpl struct {
	Service *service.Service
	UserDao dao.UserDao
}

func NewUserService(s *service.Service, userDaoImpl *dao.UserDaoImpl) UserService {
	return &UserServiceImpl{
		Service: s,
		UserDao: userDaoImpl,
	}
}
