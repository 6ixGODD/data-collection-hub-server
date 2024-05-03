package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type UserService interface {
}

type UserServiceImpl struct {
	Service *service.Core
	userDao dao.UserDao
}

func NewUserService(s *service.Core, userDao dao.UserDao) UserService {
	return &UserServiceImpl{
		Service: s,
		userDao: userDao,
	}
}
