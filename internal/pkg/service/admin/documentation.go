package admin

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DocumentationService interface {
}

type documentationServiceImpl struct {
	service          *service.Service
	documentationDao dao.DocumentationDao
}

func NewDocumentationService(s *service.Service, documentationDao dao.DocumentationDao) DocumentationService {
	return &documentationServiceImpl{
		service:          s,
		documentationDao: documentationDao,
	}
}
