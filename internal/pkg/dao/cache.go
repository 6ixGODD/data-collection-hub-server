package dao

import (
	"context"
	"errors"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/models"
	rds "data-collection-hub-server/pkg/redis"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Redis  *rds.Redis
	Config *config.Config
	Nil    CacheNil
}

type CacheNil struct{}

func (c CacheNil) Error() string {
	return "cache: nil"
}

func NewCache(redis *rds.Redis, config *config.Config) *Cache {
	return &Cache{
		Redis:  redis,
		Config: config,
		Nil:    CacheNil{},
	}
}

func (c *Cache) Get(ctx context.Context, key string) (*string, error) {
	result, err := c.Redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Cache) Set(ctx context.Context, key string, value string) error {
	err := c.Redis.RedisClient.Set(ctx, key, value, c.Config.BaseConfig.CacheTTL).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetList(ctx context.Context, key string) (*models.CacheList, error) {
	result, err := c.Redis.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var cacheList models.CacheList
	err = json.Unmarshal([]byte(result), &cacheList)
	if err != nil {
		return nil, err
	}
	return &cacheList, nil
}

func (c *Cache) SetList(ctx context.Context, key string, cacheList *models.CacheList) error {
	cacheListJSON, err := json.Marshal(cacheList)
	if err != nil {
		return err
	}
	return c.Redis.RedisClient.Set(ctx, key, cacheListJSON, c.Config.BaseConfig.CacheTTL).Err()
}

func (c *Cache) RightPush(ctx context.Context, key string, value string) error {
	return c.Redis.RedisClient.RPush(ctx, key, value).Err()
}

func (c *Cache) LeftPop(ctx context.Context, key string) (*string, error) {
	result, err := c.Redis.RedisClient.LPop(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, c.Nil
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.Redis.RedisClient.Del(ctx, key).Err()
}

func (c *Cache) Flush(ctx context.Context, prefix *string) error {
	if prefix == nil {
		return c.Redis.RedisClient.FlushAll(ctx).Err()
	}
	keys, err := c.Redis.RedisClient.Keys(ctx, *prefix+"*").Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		err := c.Redis.RedisClient.Del(ctx, key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
