package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DataAuditService interface {
}

type DataAuditServiceImpl struct {
	core               *service.Core
	instructionDataDao dao.InstructionDataDao
}

func NewDataAuditService(s *service.Core, instructionDataDao dao.InstructionDataDao) DataAuditService {
	return &DataAuditServiceImpl{
		core:               s,
		instructionDataDao: instructionDataDao,
	}
}
