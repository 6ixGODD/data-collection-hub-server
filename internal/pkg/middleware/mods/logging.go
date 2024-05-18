package mods

import (
	"time"

	"data-collection-hub-server/pkg/redis"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type LoggingMiddleware struct {
	Zap   *logging.Zap
	Redis *redis.Redis
}

func (l *LoggingMiddleware) Register(app *fiber.App) {
	app.Use(l.loggingMiddleware())
}

func (l *LoggingMiddleware) loggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			sysLogger,
			reqLogger,
			loginLogger *zap.Logger
			err      error
			sysCtx   = l.Zap.SetTagInContext(c.UserContext(), logging.SystemTag)
			reqCtx   = l.Zap.SetTagInContext(c.UserContext(), logging.RequestTag)
			loginCtx = l.Zap.SetTagInContext(c.UserContext(), logging.LoginTag)
		)
		if sysLogger, err = l.Zap.GetLogger(sysCtx); err != nil {
			return err
		}
		if reqLogger, err = l.Zap.GetLogger(reqCtx); err != nil {
			return err
		}
		if loginLogger, err = l.Zap.GetLogger(loginCtx); err != nil {
			return err
		}
		if err != nil {
			return err
		}
		err = c.Next()
		if err != nil {
			sysLogger.Error(
				"c.Next()",
				zap.Error(err),
			)
			return err
		}

		if c.Response().StatusCode() >= 400 {
			reqLogger.Warn(
				"Request",
				zap.String("request-id", c.Get("X-Request-Id")),
				zap.String("path", c.Path()),
				zap.String("method", c.Method()),
				zap.String("ip", c.IP()),
				zap.String("user-agent", c.Get("User-Agent")),
				zap.Any("query", c.Request().URI().QueryArgs()),
				zap.Any("form", c.Request().PostArgs()),
				zap.Any("body", c.Body()),
				zap.Int("status", c.Response().StatusCode()),
				zap.Any("response", c.Response().Body()),
				zap.Duration("latency", *c.Context().Value("latency").(*time.Duration)),
			)
		} else {
			reqLogger.Info(
				"Request",
				zap.String("request-id", c.Get("X-Request-Id")),
				zap.String("path", c.Path()),
				zap.String("method", c.Method()),
				zap.String("ip", c.IP()),
				zap.String("user-agent", c.Get("User-Agent")),
				zap.Any("query", c.Request().URI().QueryArgs()),
				zap.Any("form", c.Request().PostArgs()),
				zap.Any("body", c.Body()),
				zap.Int("status", c.Response().StatusCode()),
				zap.Any("response", c.Response().Body()),
				zap.Duration("latency", *c.Context().Value("latency").(*time.Duration)),
			)
		}

		if c.Path() == "/api/v1/login" && c.Response().StatusCode() < 400 {
			loginLogger.Info(
				"Login",
				zap.String("request-id", c.Get("X-Request-Id")),
				zap.String("user-id", c.Context().Value("userID").(string)),
				zap.String("ip", c.IP()),
				zap.String("user-agent", c.Get("User-Agent")),
				zap.Any("body", c.Body()),
				zap.Int("status", c.Response().StatusCode()),
				zap.Any("response", c.Response().Body()),
				zap.Duration("latency", *c.Context().Value("latency").(*time.Duration)),
			)
		}
		return nil
	}
}
