package global

import (
	"data-collection-hub-server/core/config"
	"go.uber.org/zap"
)

// TODO: Global variables

var (
	LOGGER *zap.Logger
	CONFIG *config.Config
)
