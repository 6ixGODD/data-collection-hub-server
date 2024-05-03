package common

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type ProfileService interface {
}

type profileServiceImpl struct {
	service *service.Service
	userDao dao.UserDao
}

func NewProfileService(s *service.Service, userDao dao.UserDao) ProfileService {
	return &profileServiceImpl{
		service: s,
		userDao: userDao,
	}
}
