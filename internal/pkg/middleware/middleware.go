package middleware

import (
	mods2 "data-collection-hub-server/internal/pkg/middleware/mods"
	"github.com/gofiber/fiber/v2"
)

type Middleware struct {
	AuthMiddleware       *mods2.AuthMiddleware
	LoggingMiddleware    *mods2.LoggingMiddleware
	PrometheusMiddleware *mods2.PrometheusMiddleware
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
