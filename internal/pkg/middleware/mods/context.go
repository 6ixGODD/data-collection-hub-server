package mods

import (
	"data-collection-hub-server/internal/pkg/config"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		userID, ok := c.Locals(config.UserIDKey).(primitive.ObjectID)
		if ok {
			ctx = m.Zap.SetUserIDInContext(ctx, userID.Hex())
		}
		c.SetUserContext(ctx)
		return c.Next()
	}
}
