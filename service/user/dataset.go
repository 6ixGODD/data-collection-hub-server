package user

import (
	"data-collection-hub-server/service"
)

type DatasetService interface {
}

type DatasetServiceImpl struct {
	*service.Service
}
