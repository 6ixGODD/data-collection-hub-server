package mods

import (
	"data-collection-hub-server/pkg/jwt"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	Jwt *jwt.Jwt
	Zap *logging.Zap
}

func NewAuthMiddleware(jwt *jwt.Jwt, zap *logging.Zap) *AuthMiddleware {
	return &AuthMiddleware{
		Jwt: jwt,
		Zap: zap,
	}
}

func (a *AuthMiddleware) Register(app *fiber.App) {

}

func (a *AuthMiddleware) authMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"message": "Unauthorized",
				},
			)
		}
		return c.Next()
	}
}
