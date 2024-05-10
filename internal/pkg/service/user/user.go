package user

import (
	"data-collection-hub-server/internal/pkg/service/user/mods"
)

type User struct {
	DatasetService   mods.DatasetService
	StatisticService mods.StatisticService
}

func New(datasetService mods.DatasetService, statisticService mods.StatisticService) *User {
	return &User{
		DatasetService:   datasetService,
		StatisticService: statisticService,
	}
}
