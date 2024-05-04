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
}

func NewStatisticService(s *service.Core, instructionDataDao dao.InstructionDataDao) StatisticService {
	return &StatisticServiceImpl{
		core:               s,
		instructionDataDao: instructionDataDao,
	}
}
