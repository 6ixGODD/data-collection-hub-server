package admin

import (
	dao "data-collection-hub-server/dal/modules"
	"data-collection-hub-server/service"
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
