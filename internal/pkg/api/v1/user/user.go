package user

import (
	"data-collection-hub-server/internal/pkg/api/v1/user/mods"
)

type User struct {
	*mods.DatasetApi
	*mods.StatisticApi
}

func New(datasetApi *mods.DatasetApi, statisticApi *mods.StatisticApi) User {
	return User{datasetApi, statisticApi}
}
