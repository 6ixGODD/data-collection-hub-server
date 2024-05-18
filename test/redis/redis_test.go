package redis__test

import (
	"context"
	"testing"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/redis"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	cfg, err := config.New()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	ctx := context.Background()
	r, err := redis.New(ctx, cfg.RedisConfig.GetRedisOptions())
	assert.NoError(t, err)
	assert.NotNil(t, r)
	t.Logf("redis: %+v", r)

	err = r.RedisClient.Ping(ctx).Err()
	assert.NoError(t, err)

	err = r.RedisClient.Set(ctx, "key", "value", 0).Err()
	assert.NoError(t, err)

	val, err := r.RedisClient.Get(ctx, "key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "value", val)
}
