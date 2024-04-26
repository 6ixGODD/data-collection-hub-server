package admin

import (
	dao "data-collection-hub-server/dal/modules"
	"data-collection-hub-server/service"
)

type DocumentationService interface {
}

type DocumentationServiceImpl struct {
	*service.Service
	DocumentationDao dao.DocumentationDao
}

func NewDocumentationService(s *service.Service, documentationDaoImpl *dao.DocumentationDaoImpl) DocumentationService {
	return &DocumentationServiceImpl{
		Service:          s,
		DocumentationDao: documentationDaoImpl,
	}
}
