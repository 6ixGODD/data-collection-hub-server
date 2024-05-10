package middleware

import (
	"data-collection-hub-server/pkg/middleware/mods"
)

type Middleware struct {
	AuthMiddleware       *mods.AuthMiddleware
	LoggingMiddleware    *mods.LoggingMiddleware
	PrometheusMiddleware *mods.PrometheusMiddleware
}

func New(
	authMiddleware *mods.AuthMiddleware, loggingMiddleware *mods.LoggingMiddleware,
	prometheusMiddleware *mods.PrometheusMiddleware,
) *Middleware {
	return &Middleware{
		AuthMiddleware:       authMiddleware,
		LoggingMiddleware:    loggingMiddleware,
		PrometheusMiddleware: prometheusMiddleware,
	}
}

func (m *Middleware) Register() {
	// TODO: Register middleware here.
}
