package dal

import (
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Dao struct {
	MongoDB     *qmgo.Database
	RedisClient *redis.Client
	Logger      *zap.Logger
}

func NewDao(mongoDB *qmgo.Database, redisClient *redis.Client, logger *zap.Logger) *Dao {
	return &Dao{
		MongoDB:     mongoDB,
		RedisClient: redisClient,
		Logger:      logger,
	}
}
