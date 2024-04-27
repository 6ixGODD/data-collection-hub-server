package redis

import (
	"data-collection-hub-server/internal/pkg/config/modules"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var (
	redisClient *redis.Client
	redisConfig *modules.RedisConfig
)

func InitRedis(ctx context.Context, config *modules.RedisConfig) (err error) {
	redisConfig = config
	redisClient = redis.NewClient(&redis.Options{
		Addr:            redisConfig.Addr,
		ClientName:      redisConfig.ClientName,
		DB:              redisConfig.DB,
		MaxRetries:      redisConfig.MaxRetries,
		MinRetryBackoff: redisConfig.MinRetryBackoff,
		MaxRetryBackoff: redisConfig.MaxRetryBackoff,
		DialTimeout:     redisConfig.DialTimeout,
		ReadTimeout:     redisConfig.ReadTimeout,
		WriteTimeout:    redisConfig.WriteTimeout,
		PoolSize:        redisConfig.PoolSize,
		PoolTimeout:     redisConfig.PoolTimeout,
		MinIdleConns:    redisConfig.MinIdleConns,
		MaxIdleConns:    redisConfig.MaxIdleConns,
		MaxActiveConns:  redisConfig.MaxActiveConns,
		ConnMaxIdleTime: redisConfig.ConnMaxIdleTime,
		ConnMaxLifetime: redisConfig.ConnMaxLifetime,
		// CredentialsProvider:   nil,
		// Username:              config.Username,
		// Password:              config.Password,
		// OnConnect:             nil,
	})
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func getRedisClient() (r *redis.Client, e error) {
	if redisClient == nil {
		// retry to connect
		if err := InitRedis(context.Background(), redisConfig); err != nil {
			return nil, err
		}
	}
	return redisClient, nil
}

func CloseRedis(ctx context.Context) error {
	return redisClient.Close()
}
