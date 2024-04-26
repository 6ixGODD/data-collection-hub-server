package user

import (
	dao "data-collection-hub-server/dal/modules"
	"data-collection-hub-server/service"
)

type StatisticService interface {
}

type StatisticServiceImpl struct {
	*service.Service
	instructionDataDao dao.InstructionDataDao
}

func NewStatisticService(s *service.Service, instructionDataDaoImpl *dao.InstructionDataDaoImpl) StatisticService {
	return &StatisticServiceImpl{
		Service:            s,
		instructionDataDao: instructionDataDaoImpl,
	}
}
