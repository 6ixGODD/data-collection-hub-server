package redis_test

import (
	"testing"

	"data-collection-hub-server/test/wire"
	"github.com/stretchr/testify/assert"
)

func TestRedis(t *testing.T) {
	var (
		injector = wire.GetInjector()
		r        = injector.Redis
		ctx      = injector.Ctx
		err      error
	)
	t.Logf("redis: %+v", r)
	err = r.RedisClient.Ping(ctx).Err()
	assert.NoError(t, err)

	err = r.RedisClient.Set(ctx, "key", "value", 0).Err()
	assert.NoError(t, err)

	val, err := r.RedisClient.Get(ctx, "key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "value", val)
}
