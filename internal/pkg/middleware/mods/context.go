package mods

import (
	"data-collection-hub-server/internal/pkg/config"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/gofiber/fiber/v2"
)

type ContextMiddleware struct {
	Zap *logging.Zap
}

func (m *ContextMiddleware) Register(app *fiber.App) {
	app.Use(m.contextMiddleware())
}

func (m *ContextMiddleware) contextMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		ctx = m.Zap.SetRequestIDInContext(ctx, c.Get(fiber.HeaderXRequestID))
		userID, ok := c.Locals(config.UserIDKey).(string)
		if ok {
			ctx = m.Zap.SetUserIDInContext(ctx, userID)
		}
		c.SetUserContext(ctx)
		return c.Next()
	}
}
