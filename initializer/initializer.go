package initializer

import (
	"os"

	"data-collection-hub-server/core/config"
	"data-collection-hub-server/core/zap"
	"data-collection-hub-server/global"
	"github.com/joho/godotenv"
)

// TODO: Init global variables, db, wire, etc. return a app instance

func InitApp() {
	// Load .env file if exists
	if envFile := os.Getenv("APP_ENV_FILE"); envFile != "" {
		_ = godotenv.Load(envFile)
	} else if _, err := os.Stat("../.env"); err == nil {
		_ = godotenv.Load("../.env")
	} else {
		_ = godotenv.Load()
	}

	// TODO: Init global variables
	global.CONFIG = config.GetConfig()
	global.LOGGER = zap.GetLogger()

	// TODO: Init wire
	// TODO: Init db

	// TODO: return app instance
}
