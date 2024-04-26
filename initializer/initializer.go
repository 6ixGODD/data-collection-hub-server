package initializer

import (
	"context"
	"fmt"

	"data-collection-hub-server/core/config"
	"data-collection-hub-server/core/mongo"
	"data-collection-hub-server/core/redis"
	log "data-collection-hub-server/core/zap"
	"data-collection-hub-server/hooks"
	"data-collection-hub-server/router"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewApp factory function that initializes the application and returns a fiber.App instance.
func NewApp(appConfig config.Config) (*fiber.App, error) {
	ctx := context.Background()
	// Init Zap log
	if err := log.InitLogger(&appConfig.ZapConfig); err != nil {
		return nil, err
	}
	// Init MongoDB
	if err := mongo.InitMongo(ctx, &appConfig.MongoConfig); err != nil {
		return nil, err
	}
	// Init Redis
	if err := redis.InitRedis(ctx, &appConfig.RedisConfig); err != nil {
		return nil, err
	}

	// TODO: Init wire

	app := fiber.New(fiber.Config{
		Prefork:                 appConfig.FiberConfig.Prefork,
		ServerHeader:            appConfig.FiberConfig.ServerHeader,
		BodyLimit:               appConfig.FiberConfig.BodyLimit,
		Concurrency:             appConfig.FiberConfig.Concurrency,
		ReadTimeout:             appConfig.FiberConfig.ReadTimeout,
		WriteTimeout:            appConfig.FiberConfig.WriteTimeout,
		IdleTimeout:             appConfig.FiberConfig.IdleTimeout,
		ReadBufferSize:          appConfig.FiberConfig.ReadBufferSize,
		WriteBufferSize:         appConfig.FiberConfig.WriteBufferSize,
		ProxyHeader:             appConfig.FiberConfig.ProxyHeader,
		DisableStartupMessage:   appConfig.FiberConfig.DisableStartupMessage,
		AppName:                 appConfig.BaseConfig.AppName,
		ReduceMemoryUsage:       appConfig.FiberConfig.ReduceMemoryUsage,
		EnableTrustedProxyCheck: appConfig.FiberConfig.EnableTrustedProxyCheck,
		TrustedProxies:          appConfig.FiberConfig.TrustedProxies,
		EnablePrintRoutes:       appConfig.FiberConfig.EnablePrintRoutes,
		JSONDecoder:             json.Unmarshal,
		JSONEncoder:             json.Marshal,
	})

	// Register limiter middleware
	app.Use(limiter.New(limiter.Config{
		Max:               appConfig.LimiterConfig.Max,
		Expiration:        appConfig.LimiterConfig.Expiration,
		LimiterMiddleware: limiter.SlidingWindow{},
		// LimitReached: nil,
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

	return app, nil
}

// Run function that initializes the application and starts the server.
func Run(config config.Config, appHost string, appPort string) {
	app, err := NewApp(config)
	if err != nil {
		fmt.Println("App initialization failed with error: ", err)
	}

	if appHost == "" {
		appHost = config.BaseConfig.AppHost
	}
	if appPort == "" {
		appPort = config.BaseConfig.AppPort
	}

	log.GetLogger().Fatal(
		"Server run failed",
		zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: app.Listen(":" + appPort)},
	)
}
