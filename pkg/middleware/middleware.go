package middleware

import (
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

type Middleware struct {
	Redis  *redis.Redis
	Logger *zap.Zap
}
