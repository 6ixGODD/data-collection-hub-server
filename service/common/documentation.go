package common

import (
	"data-collection-hub-server/service"
)

type DocumentationService interface {
}

type DocumentationServiceImpl struct {
	*service.Service
}
