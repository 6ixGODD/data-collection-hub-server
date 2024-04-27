package common

import (
	"data-collection-hub-server/internal/pkg/service"
)

type AuthService interface {
}

type AuthServiceImpl struct {
	*service.Service
}
