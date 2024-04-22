package admin

import (
	"data-collection-hub-server/service"
)

type UserService interface {
}

type UserServiceImpl struct {
	*service.Service
}
