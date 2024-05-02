package common

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type ProfileService interface {
}

type ProfileServiceImpl struct {
	Service *service.Service
	userDao dao.UserDao
}

func NewProfileService(s *service.Service, userDaoImpl *dao.UserDaoImpl) ProfileService {
	return &ProfileServiceImpl{
		Service: s,
		userDao: userDaoImpl,
	}
}
