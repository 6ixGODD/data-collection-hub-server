package common

import (
	"data-collection-hub-server/internal/pkg/service"
)

type DocumentationService interface {
}

type DocumentationServiceImpl struct {
	*service.Service
}
