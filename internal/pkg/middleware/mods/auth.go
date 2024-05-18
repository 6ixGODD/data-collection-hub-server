package mods

import (
	"fmt"

	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/jwt"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	Jwt *jwt.Jwt
	Zap *logging.Zap
}

func (a *AuthMiddleware) Register(app *fiber.App) {
	app.Use(a.authMiddleware())
}

func (a *AuthMiddleware) authMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return errors.TokenMissed(fmt.Errorf("token missed"))
		}
		sub, err := a.Jwt.VerifyToken(token)
		if err != nil {
			return errors.InvalidToken(err)
		}
		userCtx := a.Zap.SetUserIDWithContext(c.UserContext(), sub)
		c.SetUserContext(userCtx)
		return c.Next()
	}
}
