package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type StatisticService interface {
}

type StatisticServiceImpl struct {
	core               *service.Core
	instructionDataDao dao.InstructionDataDao
	userDao            dao.UserDao
}

func NewStatisticService(s *service.Core, instructionDataDao dao.InstructionDataDao, userDao dao.UserDao) StatisticService {
	return &StatisticServiceImpl{
		core:               s,
		instructionDataDao: instructionDataDao,
		userDao:            userDao,
	}
}
