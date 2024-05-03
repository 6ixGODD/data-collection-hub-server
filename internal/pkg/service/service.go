package service

import (
	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

// Core is the service core struct.
type Core struct {
	Logger *zap.Zap
	Redis  *redis.Redis
	Config *config.Config
	Jwt    *jwt.Jwt
}

// NewCore creates a new instance of Core.
func NewCore(logger *zap.Zap, redis *redis.Redis, config *config.Config, jwt *jwt.Jwt) *Core {
	return &Core{
		Logger: logger,
		Redis:  redis,
		Config: config,
		Jwt:    jwt,
	}
}

type Service struct {
}
