package hooks

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// Shutdown hooks to close the application gracefully
func Shutdown(ctx context.Context, app *fiber.App) error {
	// // Close Mongo
	// if err := mongo.CloseMongo(ctx); err != nil {
	// 	return err
	// }
	// // Close Redis
	// if err := redis.CloseRedis(ctx); err != nil {
	// 	return err
	// }

	// Close Fiber
	return app.Shutdown()
}
