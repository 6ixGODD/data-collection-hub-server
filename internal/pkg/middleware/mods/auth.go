package mods

import (
	e "errors"
	"fmt"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/dao"
	"data-collection-hub-server/pkg/errors"
	auth "data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/utils/check"
	"data-collection-hub-server/pkg/utils/crypt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	Jwt    *auth.Jwt
	Cache  *dao.Cache
	Config *config.Config
}

func (a *AuthMiddleware) Register(app *fiber.App) {
	app.Use(a.authMiddleware())
}

func (a *AuthMiddleware) authMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		for _, path := range a.Config.MiddlewareConfig.AuthConfig.SkippedPathPrefixes { // skip auth for some paths
			if c.Path() == path {
				return c.Next()
			}
		}
		token := c.Get(fiber.HeaderAuthorization)
		if token == "" {
			return errors.TokenMissed(fmt.Errorf("token missed"))
		}
		if !check.IsBearerToken(token) {
			return errors.TokenInvalid(fmt.Errorf("token should be bearer token (start with 'Bearer ' or 'bearer ')"))
		}
		token = token[7:] // remove 'Bearer '
		if ok, err := a.Cache.Get(c.Context(), crypt.MD5(token)); err == nil && *ok == config.CacheTrue {
			return errors.TokenInvalid(fmt.Errorf("token has been revoked"))
		}
		sub, err := a.Jwt.VerifyToken(token)
		if err != nil {
			var ve *jwt.ValidationError
			if e.As(err, &ve) {
				if ve.Errors == jwt.ValidationErrorExpired {
					return errors.TokenExpired(fmt.Errorf("token is expired"))
				}
			}
			return errors.TokenInvalid(fmt.Errorf("token invalid"))
		}
		c.Locals(config.UserIDKey, sub)
		return c.Next()
	}
}
