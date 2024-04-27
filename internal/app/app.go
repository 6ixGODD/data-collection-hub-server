package app

import (
	"context"
	"fmt"
	"os"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/hooks"
	"data-collection-hub-server/internal/pkg/router"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	log "data-collection-hub-server/pkg/zap"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewApp factory function that initializes the application and returns a fiber.App instance.
func NewApp(appConfig config.Config, ctx context.Context) (*fiber.App, error) {
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
func Run(config config.Config, appHost string, appPort string, ctx context.Context) {
	app, err := NewApp(config, ctx)
	if err != nil {
		fmt.Println("App initialization failed with error: ", err)
	}
	logger := log.GetLoggerWithContext(ctx)

	if appHost == "" {
		appHost = config.BaseConfig.AppHost
	}
	if appPort == "" {
		appPort = config.BaseConfig.AppPort
	}
	logger.Info(
		"Service is starting",
		zap.Field{Key: "host", Type: zapcore.StringType, Interface: appHost},
		zap.Field{Key: "port", Type: zapcore.StringType, Interface: appPort},
		zap.Field{Key: "version", Type: zapcore.StringType, Interface: config.BaseConfig.AppVersion},
		zap.Field{Key: "pid", Type: zapcore.Int64Type, Interface: os.Getpid()},
	)
	if config.BaseConfig.EnableTls {
		logger.Fatal(
			"Server run failed",
			zap.Field{
				Key:  "error",
				Type: zapcore.ErrorType,
				Interface: app.ListenTLS(
					appHost+":"+appPort, config.BaseConfig.TlsCertFile, config.BaseConfig.TlsKeyFile,
				),
			},
		)
	} else {
		logger.Fatal(
			"Server run failed",
			zap.Field{Key: "error", Type: zapcore.ErrorType, Interface: app.Listen(appHost + ":" + appPort)},
		)
	}
}
