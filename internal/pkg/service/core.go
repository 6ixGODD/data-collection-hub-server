package service

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	logging "data-collection-hub-server/pkg/zap"
	"go.uber.org/zap"
)

// Core contains the core components of the service.
type Core struct {
	Config *config.Config
	Logger *zap.Logger
}

func NewCore(ctx context.Context, config *config.Config, zap *logging.Zap) (*Core, error) {
	c := zap.SetTagInContext(ctx, logging.ServiceTag)
	logger, err := zap.GetLogger(c)
	if err != nil {
		return nil, err
	}
	return &Core{
		Config: config,
		Logger: logger,
	}, nil
}
