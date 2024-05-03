package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DatasetService interface {
}

type DatasetServiceImpl struct {
	service            *service.Core
	instructionDataDao dao.InstructionDataDao
	operationLogDao    dao.OperationLogDao
}

func NewDatasetService(s *service.Core, instructionDataDao dao.InstructionDataDao, operationLogDao dao.OperationLogDao) DatasetService {
	return &DatasetServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
		operationLogDao:    operationLogDao,
	}
}
