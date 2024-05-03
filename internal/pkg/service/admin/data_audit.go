package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DataAuditService interface {
}

type dataAuditServiceImpl struct {
	service            *service.Service
	instructionDataDao dao.InstructionDataDao
}

func NewDataAuditService(s *service.Service, instructionDataDao dao.InstructionDataDao) DataAuditService {
	return &dataAuditServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
	}
}
