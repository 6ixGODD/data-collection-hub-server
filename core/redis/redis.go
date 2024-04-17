package redis

import (
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func InitRedis() {
}

func getClient() *redis.Client {
	return client
}
