package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type StatisticService interface {
}

type StatisticServiceImpl struct {
	*service.Service
	dao.InstructionDataDao
	dao.UserDao
}

func NewStatisticService(s *service.Service, instructionDataDaoImpl *dao.InstructionDataDaoImpl, userDaoImpl *dao.UserDaoImpl) StatisticService {
	return &StatisticServiceImpl{
		Service:            s,
		InstructionDataDao: instructionDataDaoImpl,
		UserDao:            userDaoImpl,
	}
}
