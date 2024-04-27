package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DataAuditService interface {
}

type DataAuditServiceImpl struct {
	*service.Service
	InstructionDataDao dao.InstructionDataDao
}

func NewDataAuditService(s *service.Service, instructionDataDaoImpl *dao.InstructionDataDaoImpl) DataAuditService {
	return &DataAuditServiceImpl{
		Service:            s,
		InstructionDataDao: instructionDataDaoImpl,
	}
}
