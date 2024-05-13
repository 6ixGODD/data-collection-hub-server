package app

import (
	"context"
	"os"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/errors"
	"data-collection-hub-server/internal/pkg/hooks"
	"data-collection-hub-server/internal/pkg/router"
	"data-collection-hub-server/pkg/cron"
	"data-collection-hub-server/pkg/middleware"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/casbin/mongodb-adapter/v3"
	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

type App struct {
	app        *fiber.App
	logger     *logging.Zap
	config     *config.Config
	router     *router.Router
	middleware *middleware.Middleware
	Scheduler  *cron.Scheduler
	ctx        context.Context
}

// New factory function that initializes the application and returns a fiber.App instance.
func New(
	ctx context.Context, logger *logging.Zap, config *config.Config, router *router.Router,
	middleware *middleware.Middleware, scheduler *cron.Scheduler,
) (*App, error) {
	app := &App{
		logger:     logger,
		config:     config,
		router:     router,
		middleware: middleware,
		Scheduler:  scheduler,
		ctx:        ctx,
	}

	if err := app.Init(); err != nil {
		return nil, err
	}
	return app, nil
}

func (a *App) Init() error {
	app := fiber.New(
		fiber.Config{
			Prefork:                 a.config.FiberConfig.Prefork,
			ServerHeader:            a.config.FiberConfig.ServerHeader,
			BodyLimit:               a.config.FiberConfig.BodyLimit,
			Concurrency:             a.config.FiberConfig.Concurrency,
			ReadTimeout:             a.config.FiberConfig.ReadTimeout,
			WriteTimeout:            a.config.FiberConfig.WriteTimeout,
			IdleTimeout:             a.config.FiberConfig.IdleTimeout,
			ReadBufferSize:          a.config.FiberConfig.ReadBufferSize,
			WriteBufferSize:         a.config.FiberConfig.WriteBufferSize,
			ProxyHeader:             a.config.FiberConfig.ProxyHeader,
			DisableStartupMessage:   a.config.FiberConfig.DisableStartupMessage,
			AppName:                 a.config.BaseConfig.AppName,
			ReduceMemoryUsage:       a.config.FiberConfig.ReduceMemoryUsage,
			EnableTrustedProxyCheck: a.config.FiberConfig.EnableTrustedProxyCheck,
			TrustedProxies:          a.config.FiberConfig.TrustedProxies,
			EnablePrintRoutes:       a.config.FiberConfig.EnablePrintRoutes,
			ErrorHandler:            errors.ErrorHandler,
			JSONDecoder:             json.Unmarshal, // Use go-json for enhanced JSON decoding performance
			JSONEncoder:             json.Marshal,
		},
	)

	// Register middleware
	// Register limiter middleware
	app.Use(
		limiter.New(
			limiter.Config{
				Max:               a.config.LimiterConfig.Max,
				Expiration:        a.config.LimiterConfig.Expiration,
				LimiterMiddleware: limiter.SlidingWindow{},
				// LimitReached: nil,
			},
		),
	)

	// Register cors middleware
	app.Use(
		cors.New(
			cors.Config{
				// Next:             nil,
				AllowOrigins:     a.config.CorsConfig.AllowOrigins,
				AllowMethods:     a.config.CorsConfig.AllowMethods,
				AllowHeaders:     a.config.CorsConfig.AllowHeaders,
				AllowCredentials: a.config.CorsConfig.AllowCredentials,
				ExposeHeaders:    a.config.CorsConfig.ExposeHeaders,
				MaxAge:           a.config.CorsConfig.MaxAge,
			},
		),
	)

	// Register request id middleware
	app.Use(requestid.New())

	// Register logger middleware

	// TODO: Add more middleware here

	// Register hooks
	app.Hooks().OnShutdown(
		func() error {
			return hooks.Shutdown(a.ctx, app)
		},
	)

	// TODO: Add more hooks here

	// TODO: Add scheduler tasks

	// Ping
	app.Get(
		"/ping", func(c *fiber.Ctx) error {
			return c.SendString("pong")
		},
	)

	// Register routers
	adapter, err := mongodbadapter.NewAdapter(a.config.CasbinConfig.PolicyAdapterUrl)
	if err != nil {
		return err
	}
	rbac := casbin.New(
		casbin.Config{
			ModelFilePath: a.config.CasbinConfig.ModelPath,
			PolicyAdapter: adapter,
			Lookup: func(c *fiber.Ctx) string {
				return c.Locals(config.KeyUserID).(string)
			},
		},
	)
	a.router.RegisterRouter(app, rbac)

	a.app = app

	return nil
}

// Run function starts the application server.
func (a *App) Run(addr string, enableTls bool) {
	a.logger.SetTagInContext(a.ctx, logging.SystemTag)
	a.logger.Logger.Info(
		"Server is starting",
		zap.String("Addr", addr),
		zap.String("version", a.config.BaseConfig.AppVersion),
		zap.Int64("pid", int64(os.Getpid())),
	)

	if a.config.BaseConfig.EnableTls || enableTls {
		if err := a.app.ListenTLS(addr, a.config.BaseConfig.TlsCertFile, a.config.BaseConfig.TlsKeyFile); err != nil {
			a.logger.Logger.Fatal("Server run failed", zap.Error(err))
		}
	} else {
		if err := a.app.Listen(addr); err != nil {
			a.logger.Logger.Fatal("Server run failed", zap.Error(err))
		}
	}
}
