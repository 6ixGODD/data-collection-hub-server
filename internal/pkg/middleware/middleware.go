package middleware

import (
	ware "data-collection-hub-server/internal/pkg/middleware/mods"
	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	AuthMiddleware       *ware.AuthMiddleware
	LoggingMiddleware    *ware.LoggingMiddleware
	PrometheusMiddleware *ware.PrometheusMiddleware
}

func (m *Middleware) Register(app *fiber.App) error {
	m.AuthMiddleware.Register(app)
	m.LoggingMiddleware.Register(app)
	err := m.PrometheusMiddleware.Register(app)
	if err != nil {
		return err
	}
	return nil
}
