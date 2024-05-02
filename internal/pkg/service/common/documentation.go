package common

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DocumentationService interface {
}

type DocumentationServiceImpl struct {
	Service          *service.Service
	documentationDao dao.DocumentationDao
}

func NewDocumentationService(s *service.Service, documentationDaoImpl *dao.DocumentationDaoImpl) DocumentationService {
	return &DocumentationServiceImpl{
		Service:          s,
		documentationDao: documentationDaoImpl,
	}
}
