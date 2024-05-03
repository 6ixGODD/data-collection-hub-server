package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type UserService interface {
}

type userServiceImpl struct {
	Service *service.Service
	userDao dao.UserDao
}

func NewUserService(s *service.Service, userDao dao.UserDao) UserService {
	return &userServiceImpl{
		Service: s,
		userDao: userDao,
	}
}
