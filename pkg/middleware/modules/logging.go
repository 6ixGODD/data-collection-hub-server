package modules

import (
	"data-collection-hub-server/pkg/middleware"
)

type LoggingMiddleware struct {
	Core *middleware.Core
}
