package modules

import (
	dao "data-collection-hub-server/internal/pkg/dal/modules"
	"data-collection-hub-server/internal/pkg/service"
)

type DocumentationService interface {
}

type DocumentationServiceImpl struct {
	core             *service.Core
	documentationDao dao.DocumentationDao
}

func NewDocumentationService(s *service.Core, documentationDao dao.DocumentationDao) DocumentationService {
	return &DocumentationServiceImpl{
		core:             s,
		documentationDao: documentationDao,
	}
}
