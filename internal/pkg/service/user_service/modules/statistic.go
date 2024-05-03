package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type StatisticService interface {
}

type StatisticServiceImpl struct {
	service            *service.Core
	instructionDataDao dao.InstructionDataDao
}

func NewStatisticService(s *service.Core, instructionDataDao dao.InstructionDataDao) StatisticService {
	return &StatisticServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
	}
}
