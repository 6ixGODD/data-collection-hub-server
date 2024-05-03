package user

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type StatisticService interface {
}

type statisticServiceImpl struct {
	service            *service.Service
	instructionDataDao dao.InstructionDataDao
}

func NewStatisticService(s *service.Service, instructionDataDao dao.InstructionDataDao) StatisticService {
	return &statisticServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
	}
}
