package modules

import (
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/middleware"
)

type AuthMiddleware struct {
	Core *middleware.Core
	Jwt  *jwt.Jwt
}
