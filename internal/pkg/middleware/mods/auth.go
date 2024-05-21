package mods

import (
	"fmt"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	Jwt    *jwt.Jwt
	Config *config.Config
}

func (a *AuthMiddleware) Register(app *fiber.App) {
	app.Use(a.authMiddleware())
}

func (a *AuthMiddleware) authMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, path := range a.Config.MiddlewareConfig.AuthConfig.SkippedPathPrefixes {
			if c.Path() == path {
				return c.Next()
			}
		}
		token := c.Get(fiber.HeaderAuthorization)
		if token == "" {
			return errors.TokenMissed(fmt.Errorf("token missed"))
		}
		sub, err := a.Jwt.VerifyToken(token)
		if err != nil {
			return errors.InvalidToken(err)
		}
		c.Locals(config.UserIDKey, sub)
		return c.Next()
	}
}
