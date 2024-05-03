package dal

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

type Dao struct {
	MongoClient *mongo.Mongo
	RedisClient *redis.Redis
	Logger      *zap.Zap
	Config      *config.Config
}

func NewDao(ctx context.Context, mongoDB *mongo.Mongo, redisClient *redis.Redis, logger *zap.Zap, config config.Config) *Dao {
	logger.SetTagInContext(ctx, zap.MongoTag)
	return &Dao{
		MongoClient: mongoDB,
		RedisClient: redisClient,
		Logger:      logger,
		Config:      &config,
	}
}
