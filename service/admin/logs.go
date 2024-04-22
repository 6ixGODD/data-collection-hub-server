package admin

import (
	"data-collection-hub-server/service"
)

type LogsService interface {
}

type LogsServiceImpl struct {
	*service.Service
}
