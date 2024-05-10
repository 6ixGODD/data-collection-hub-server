package mods

import (
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

type LoggingMiddleware struct {
	Zap   *zap.Zap
	Redis *redis.Redis
}

func NewLoggingMiddleware(zap *zap.Zap, redis *redis.Redis) *LoggingMiddleware {
	return &LoggingMiddleware{
		Zap:   zap,
		Redis: redis,
	}
}
