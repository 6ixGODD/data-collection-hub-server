package app

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/domain/vo"
	"data-collection-hub-server/internal/pkg/errors"
	"data-collection-hub-server/internal/pkg/middleware"
	"data-collection-hub-server/internal/pkg/router"
	"data-collection-hub-server/internal/pkg/tasks"
	e "data-collection-hub-server/pkg/errors"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/casbin/mongodb-adapter/v3"
	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type App struct {
	App        *fiber.App
	Zap        *logging.Zap
	Logger     *zap.Logger
	Config     *config.Config
	Router     *router.Router
	Middleware *middleware.Middleware
	Tasks      *tasks.Tasks
	Mongo      *mongo.Mongo
	Redis      *redis.Redis
	Ctx        context.Context
}

// New factory function that initializes the application and returns a fiber.App instance.
func New(
	ctx context.Context, zap *logging.Zap, config *config.Config, router *router.Router,
	middleware *middleware.Middleware, tasks *tasks.Tasks, mongo *mongo.Mongo, redis *redis.Redis,
) (*App, error) {
	app := &App{
		Zap:        zap,
		Config:     config,
		Router:     router,
		Middleware: middleware,
		Tasks:      tasks,
		Mongo:      mongo,
		Redis:      redis,
		Ctx:        ctx,
	}

	if err := app.Init(ctx); err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) Init(ctx context.Context) error {
	// Set logger
	ctx = a.Zap.SetTagInContext(ctx, logging.SystemTag)
	logger, err := a.Zap.GetLogger(ctx)
	if err != nil {
		return err
	}
	a.Logger = logger

	// Set Fiber app
	app := fiber.New(
		fiber.Config{
			Prefork:                 a.Config.FiberConfig.Prefork,
			ServerHeader:            a.Config.FiberConfig.ServerHeader,
			BodyLimit:               a.Config.FiberConfig.BodyLimit,
			Concurrency:             a.Config.FiberConfig.Concurrency,
			ReadTimeout:             a.Config.FiberConfig.ReadTimeout,
			WriteTimeout:            a.Config.FiberConfig.WriteTimeout,
			IdleTimeout:             a.Config.FiberConfig.IdleTimeout,
			ReadBufferSize:          a.Config.FiberConfig.ReadBufferSize,
			WriteBufferSize:         a.Config.FiberConfig.WriteBufferSize,
			ProxyHeader:             a.Config.FiberConfig.ProxyHeader,
			DisableStartupMessage:   a.Config.FiberConfig.DisableStartupMessage,
			AppName:                 a.Config.BaseConfig.AppName,
			ReduceMemoryUsage:       a.Config.FiberConfig.ReduceMemoryUsage,
			EnableTrustedProxyCheck: a.Config.FiberConfig.EnableTrustedProxyCheck,
			TrustedProxies:          a.Config.FiberConfig.TrustedProxies,
			EnablePrintRoutes:       a.Config.FiberConfig.EnablePrintRoutes,
			ErrorHandler:            errors.ErrorHandler, // Custom error handler
			JSONDecoder:             json.Unmarshal,      // Use go-json for enhanced JSON decoding performance
			JSONEncoder:             json.Marshal,
		},
	)

	// Register Middleware
	if err := a.Middleware.Register(app); err != nil {
		return err
	}

	// Set hooks
	app.Hooks().OnShutdown(
		func() error {
			return ShutdownHandler(a.Ctx, a)
		},
	)

	// Ping

	// Set Casbin
	adapter, err := mongodbadapter.NewAdapter(a.Config.CasbinConfig.PolicyAdapterUrl)
	if err != nil {
		return err
	}
	c := casbin.New(
		casbin.Config{
			ModelFilePath: a.Config.CasbinConfig.ModelPath,
			PolicyAdapter: adapter,
			Lookup: func(c *fiber.Ctx) string {
				return c.Locals(config.UserIDKey).(string)
			},
			Forbidden: func(ctx *fiber.Ctx) error {
				return ctx.Status(fiber.StatusForbidden).JSON(
					vo.Response{
						Code:    e.CodePermissionDeny,
						Message: "Forbidden access",
						Data:    nil,
					},
				)
			},
		},
	)

	idempotencyMiddleware := a.Middleware.IdempotencyMiddleware.IdempotencyMiddleware()

	// Register routers
	a.Router.RegisterRouter(app, c, idempotencyMiddleware)

	// Set app
	a.App = app

	// Start scheduled tasks
	if err := a.Tasks.Start(); err != nil {
		return err
	}
	return nil
}
