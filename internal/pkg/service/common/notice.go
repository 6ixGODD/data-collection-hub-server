package common

import (
	"data-collection-hub-server/internal/pkg/service"
)

type NoticeService interface {
}

type NoticeServiceImpl struct {
	*service.Service
}
