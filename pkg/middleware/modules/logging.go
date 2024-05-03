package modules

import (
	"data-collection-hub-server/pkg/middleware"
)

type LoggingMiddleware struct {
	Middleware *middleware.Middleware
}
