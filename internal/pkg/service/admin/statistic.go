package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type StatisticService interface {
}

type statisticServiceImpl struct {
	service            *service.Service
	instructionDataDao dao.InstructionDataDao
	userDao            dao.UserDao
}

func NewStatisticService(s *service.Service, instructionDataDao dao.InstructionDataDao, userDao dao.UserDao) StatisticService {
	return &statisticServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
		userDao:            userDao,
	}
}
