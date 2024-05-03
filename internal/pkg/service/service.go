package service

import (
	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

// Service is the interface for the service layer.
type Service struct {
	Logger *zap.Zap
	Redis  *redis.Redis
	Config *config.Config
	Jwt    *jwt.Jwt
}

// NewService creates a new instance of Service.
func NewService(logger *zap.Zap, redis *redis.Redis, config *config.Config, jwt *jwt.Jwt) *Service {
	return &Service{
		Logger: logger,
		Redis:  redis,
		Config: config,
		Jwt:    jwt,
	}
}
