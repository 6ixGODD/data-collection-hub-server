package common

import (
	"data-collection-hub-server/service"
)

type AuthService interface {
}

type AuthServiceImpl struct {
	*service.Service
}
