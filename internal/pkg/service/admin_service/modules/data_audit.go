package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DataAuditService interface {
}

type DataAuditServiceImpl struct {
	service            *service.Core
	instructionDataDao dao.InstructionDataDao
}

func NewDataAuditService(s *service.Core, instructionDataDao dao.InstructionDataDao) DataAuditService {
	return &DataAuditServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
	}
}
