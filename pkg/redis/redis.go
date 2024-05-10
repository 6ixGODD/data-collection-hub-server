package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type Redis struct {
	RedisClient *redis.Client
	redisConfig *Config
	mutex       sync.Mutex
	ctx         context.Context
}

type Config struct {
	redisOptions *redis.Options
}

var redisInstance *Redis // Singleton

func New(ctx context.Context, options *redis.Options) (r *Redis, err error) {
	if redisInstance != nil {
		return redisInstance, nil
	}
	r = &Redis{redisConfig: &Config{options}, ctx: ctx}
	if err := r.Init(); err != nil {
		return nil, err
	}
	redisInstance = r
	return r, nil
}

func (r *Redis) Init() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.RedisClient != nil {
		return nil
	}
	client := redis.NewClient(r.redisConfig.redisOptions)
	if _, err := client.Ping(r.ctx).Result(); err != nil {
		return err
	}
	r.RedisClient = client
	return nil
}

func (r *Redis) GetClient() (client *redis.Client, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.RedisClient == nil {
		if err = r.Init(); err != nil {
			return nil, err
		}
	}
	return r.RedisClient, nil
}

func (r *Redis) Close() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.RedisClient == nil {
		return nil
	}
	err := r.RedisClient.Close()
	r.RedisClient = nil
	return err
}
