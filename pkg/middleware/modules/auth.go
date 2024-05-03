package modules

import (
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/middleware"
)

type AuthMiddleware struct {
	Middleware *middleware.Middleware
	Jwt        *jwt.Jwt
}
