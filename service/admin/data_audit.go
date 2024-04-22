package admin

import (
	"data-collection-hub-server/service"
)

type DataAuditService interface {
}

type DataAuditServiceImpl struct {
	*service.Service
}
