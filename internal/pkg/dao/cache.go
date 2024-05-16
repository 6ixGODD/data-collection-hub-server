package dao

import (
	"context"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/internal/pkg/models"
	"data-collection-hub-server/pkg/redis"
	"github.com/goccy/go-json"
)

type Cache struct {
	Redis  *redis.Redis
	Config *config.Config
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
	err = c.Redis.RedisClient.Set(ctx, key, cacheListJSON, c.Config.BaseConfig.CacheTTL).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	err := c.Redis.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
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
