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
