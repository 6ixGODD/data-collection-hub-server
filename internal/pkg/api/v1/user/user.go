package user

import (
	"data-collection-hub-server/internal/pkg/api/v1/user/mods"
)

type User struct {
	*mods.DatasetApi
	*mods.StatisticApi
}
