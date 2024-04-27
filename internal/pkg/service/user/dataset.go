package user

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DatasetService interface {
}

type DatasetServiceImpl struct {
	*service.Service
	instructionDataDao dao.InstructionDataDao
}

func NewDatasetService(s *service.Service, instructionDataDaoImpl *dao.InstructionDataDaoImpl) DatasetService {
	return &DatasetServiceImpl{
		Service:            s,
		instructionDataDao: instructionDataDaoImpl,
	}
}
