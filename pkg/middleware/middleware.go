package middleware

import (
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

type Core struct {
	Redis  *redis.Redis
	Logger *zap.Zap
}
