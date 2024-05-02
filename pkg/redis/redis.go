package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type Cache struct {
	RedisClient  *redis.Client
	RedisOptions *redis.Options
	Mutex        sync.Mutex
}

func New(ctx context.Context, options *redis.Options) (c *Cache, err error) {
	c = &Cache{RedisOptions: options}
	if err := c.Init(ctx); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Cache) Init(ctx context.Context) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.RedisClient != nil {
		return nil
	}
	client := redis.NewClient(c.RedisOptions)
	if _, err := client.Ping(ctx).Result(); err != nil {
		return err
	}
	c.RedisClient = client
	return nil
}

func (c *Cache) GetClient(ctx context.Context) (client *redis.Client, err error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.RedisClient == nil {
		if err = c.Init(ctx); err != nil {
			return nil, err
		}
	}
	return c.RedisClient, nil
}

func (c *Cache) Close() error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.RedisClient == nil {
		return nil
	}
	err := c.RedisClient.Close()
	c.RedisClient = nil
	return err
}
