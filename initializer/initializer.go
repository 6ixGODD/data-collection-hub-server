package initializer

import (
	"data-collection-hub-server/core/config"
	"data-collection-hub-server/global"
)

// TODO: Init global variables, db, wire, etc. return a app instance

func InitApp() {
	// TODO: Init global variables
	global.CONFIG = config.GetConfig()
	// TODO: Init wire
	// TODO: Init db

	// TODO: return app instance
}
