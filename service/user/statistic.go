package user

import (
	"data-collection-hub-server/service"
)

type StatisticService interface {
}

type StatisticServiceImpl struct {
	*service.Service
}
