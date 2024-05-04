package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type ProfileService interface {
}

type ProfileServiceImpl struct {
	core    *service.Core
	userDao dao.UserDao
}

func NewProfileService(s *service.Core, userDao dao.UserDao) ProfileService {
	return &ProfileServiceImpl{
		core:    s,
		userDao: userDao,
	}
}
