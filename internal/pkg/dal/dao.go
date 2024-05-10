package dal

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

type Dao struct {
	Mongo  *mongo.Mongo
	Redis  *redis.Redis
	Zap    *zap.Zap
	Config *config.Config
}

func New(ctx context.Context, mongo *mongo.Mongo, redis *redis.Redis, logger *zap.Zap, config config.Config) *Dao {
	logger.SetTagInContext(ctx, zap.MongoTag)
	return &Dao{
		Mongo:  mongo,
		Redis:  redis,
		Zap:    logger,
		Config: &config,
	}
}
