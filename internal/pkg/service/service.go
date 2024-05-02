package service

import (
	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

// Service is the interface for the service layer.
type Service struct {
	Logger *zap.Logger
	Redis  *redis.Cache
	Config *config.Config
	Jwt    *jwt.Auth
}

// NewService creates a new instance of Service.
func NewService(logger *zap.Logger, redis *redis.Cache, config *config.Config, jwt *jwt.Auth) *Service {
	return &Service{
		Logger: logger,
		Redis:  redis,
		Config: config,
		Jwt:    jwt,
	}
}
