package service

import (
	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

// Service is the service core struct.
type Service struct {
	Logger *zap.Zap
	Redis  *redis.Redis
	Config *config.Config
}

// New creates a new instance of Service.
func New(logger *zap.Zap, redis *redis.Redis, config *config.Config) *Service {
	return &Service{
		Logger: logger,
		Redis:  redis,
		Config: config,
	}
}
