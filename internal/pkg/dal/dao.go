package dal

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

type Core struct {
	MongoClient *mongo.Mongo
	RedisClient *redis.Redis
	Logger      *zap.Zap
	Config      *config.Config
}

func NewCore(ctx context.Context, mongoDB *mongo.Mongo, redisClient *redis.Redis, logger *zap.Zap, config config.Config) *Core {
	logger.SetTagInContext(ctx, zap.MongoTag)
	return &Core{
		MongoClient: mongoDB,
		RedisClient: redisClient,
		Logger:      logger,
		Config:      &config,
	}
}

type Dao struct {
}
