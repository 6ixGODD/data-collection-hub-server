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
	Ctx          context.Context
}

func New(ctx context.Context, options *redis.Options) (c *Cache, err error) {
	c = &Cache{RedisOptions: options, Ctx: ctx}
	if err := c.Init(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Cache) Init() error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.RedisClient != nil {
		return nil
	}
	client := redis.NewClient(c.RedisOptions)
	if _, err := client.Ping(c.Ctx).Result(); err != nil {
		return err
	}
	c.RedisClient = client
	return nil
}

func (c *Cache) GetClient() (client *redis.Client, err error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.RedisClient == nil {
		if err = c.Init(); err != nil {
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
