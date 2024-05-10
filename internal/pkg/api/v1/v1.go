package api

import (
	"data-collection-hub-server/internal/pkg/api/v1/admin"
	"data-collection-hub-server/internal/pkg/api/v1/common"
	"data-collection-hub-server/internal/pkg/api/v1/user"
)

type Api struct {
	AdminApi  *admin.Admin
	CommonApi *common.Common
	UserApi   *user.User
}

func New(adminApi *admin.Admin, commonApi *common.Common, userApi *user.User) *Api {
	return &Api{
		AdminApi:  adminApi,
		CommonApi: commonApi,
		UserApi:   userApi,
	}
}
