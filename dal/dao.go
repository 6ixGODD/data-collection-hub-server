package dal

import (
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Dao struct {
	mongoDB     *qmgo.Database
	redisClient *redis.Client
	logger      *zap.Logger
}
