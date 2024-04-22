package hooks

import (
	"context"

	"data-collection-hub-server/core/mongo"
	"github.com/gofiber/fiber/v2"
)

// Shutdown hooks to close the application gracefully
func Shutdown(ctx context.Context, app *fiber.App) error {
	// Close Mongo
	err := mongo.CloseMongo(ctx)
	if err != nil {
		return err
	}
	// Close Fiber
	return app.Shutdown()
}
