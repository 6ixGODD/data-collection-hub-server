package user

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DatasetService interface {
}

type DatasetServiceImpl struct {
	Service            *service.Service
	instructionDataDao dao.InstructionDataDao
	OperationLogDao    dao.OperationLogDao
}

func NewDatasetService(s *service.Service, instructionDataDaoImpl *dao.InstructionDataDaoImpl, operationLogDaoImpl *dao.OperationLogDaoImpl) DatasetService {
	return &DatasetServiceImpl{
		Service:            s,
		instructionDataDao: instructionDataDaoImpl,
		OperationLogDao:    operationLogDaoImpl,
	}
}
