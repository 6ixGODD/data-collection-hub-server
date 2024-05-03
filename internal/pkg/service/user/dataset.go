package user

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DatasetService interface {
}

type datasetServiceImpl struct {
	service            *service.Service
	instructionDataDao dao.InstructionDataDao
	operationLogDao    dao.OperationLogDao
}

func NewDatasetService(s *service.Service, instructionDataDao dao.InstructionDataDao, operationLogDao dao.OperationLogDao) DatasetService {
	return &datasetServiceImpl{
		service:            s,
		instructionDataDao: instructionDataDao,
		operationLogDao:    operationLogDao,
	}
}
