package app

import (
	"context"
	"os"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/hooks"
	"data-collection-hub-server/internal/pkg/router"
	logging "data-collection-hub-server/pkg/zap"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"go.uber.org/zap"
)

type App struct {
	app    *fiber.App
	Logger *logging.Logger
	config *config.Config
}

// NewApp factory function that initializes the application and returns a fiber.App instance.
func NewApp(config config.Config, ctx context.Context) (a *App, err error) {
	app := fiber.New(fiber.Config{
		Prefork:                 config.FiberConfig.Prefork,
		ServerHeader:            config.FiberConfig.ServerHeader,
		BodyLimit:               config.FiberConfig.BodyLimit,
		Concurrency:             config.FiberConfig.Concurrency,
		ReadTimeout:             config.FiberConfig.ReadTimeout,
		WriteTimeout:            config.FiberConfig.WriteTimeout,
		IdleTimeout:             config.FiberConfig.IdleTimeout,
		ReadBufferSize:          config.FiberConfig.ReadBufferSize,
		WriteBufferSize:         config.FiberConfig.WriteBufferSize,
		ProxyHeader:             config.FiberConfig.ProxyHeader,
		DisableStartupMessage:   config.FiberConfig.DisableStartupMessage,
		AppName:                 config.BaseConfig.AppName,
		ReduceMemoryUsage:       config.FiberConfig.ReduceMemoryUsage,
		EnableTrustedProxyCheck: config.FiberConfig.EnableTrustedProxyCheck,
		TrustedProxies:          config.FiberConfig.TrustedProxies,
		EnablePrintRoutes:       config.FiberConfig.EnablePrintRoutes,
		JSONDecoder:             json.Unmarshal, // Use go-json for enhanced JSON decoding
		JSONEncoder:             json.Marshal,
	})

	// Register limiter middleware
	app.Use(limiter.New(limiter.Config{
		Max:               config.LimiterConfig.Max,
		Expiration:        config.LimiterConfig.Expiration,
		LimiterMiddleware: limiter.SlidingWindow{},
		// LimitReached: nil,
	}))

	// Register cors middleware
	app.Use(cors.New(cors.Config{
		// Next:             nil,
		// AllowOriginsFunc: nil,
		AllowOrigins:     config.CorsConfig.AllowOrigins,
		AllowMethods:     config.CorsConfig.AllowMethods,
		AllowHeaders:     config.CorsConfig.AllowHeaders,
		AllowCredentials: config.CorsConfig.AllowCredentials,
		ExposeHeaders:    config.CorsConfig.ExposeHeaders,
		MaxAge:           config.CorsConfig.MaxAge,
	}))

	// Register hooks
	app.Hooks().OnShutdown(
		func() error {
			return hooks.Shutdown(ctx, app)
		},
	)

	// Ping
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// Register routers
	router.RegisterRouter(app)

	logger, err := logging.New(config.ZapConfig.ToZapConfig())
	if err != nil {
		return nil, err
	}

	return &App{
		app:    app,
		config: &config,
		Logger: logger,
	}, nil
}

// Run function starts the application server.
func (a *App) Run(addr string, ctx context.Context) {
	a.Logger.SetTagInContext(ctx, logging.SystemTag)

	a.Logger.Logger.Info("Server is starting",
		zap.String("Addr", addr),
		zap.String("version", a.config.BaseConfig.AppVersion),
		zap.Int64("pid", int64(os.Getpid())),
	)

	if a.config.BaseConfig.EnableTls {
		if err := a.app.ListenTLS(addr, a.config.BaseConfig.TlsCertFile, a.config.BaseConfig.TlsKeyFile); err != nil {
			a.Logger.Logger.Fatal("Server run failed", zap.Error(err))
		}
	} else {
		if err := a.app.Listen(addr); err != nil {
			a.Logger.Logger.Fatal("Server run failed", zap.Error(err))
		}
	}
}
