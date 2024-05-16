package user

import (
	"data-collection-hub-server/internal/pkg/service/user/mods"
)

type User struct {
	DatasetService   mods.DatasetService
	StatisticService mods.StatisticService
}
