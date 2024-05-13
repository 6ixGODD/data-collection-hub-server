package wire

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"data-collection-hub-server/internal/pkg/config"
	"data-collection-hub-server/pkg/jwt"
	"data-collection-hub-server/pkg/mongo"
	"data-collection-hub-server/pkg/prometheus"
	"data-collection-hub-server/pkg/redis"
	"data-collection-hub-server/pkg/zap"
)

func InitializeMongo(ctx context.Context, config *config.Config) (*mongo.Mongo, error) {
	m, err := mongo.New(
		ctx, config.MongoConfig.GetQmgoConfig(), config.MongoConfig.PingTimeoutS, config.MongoConfig.Database,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func InitializeRedis(ctx context.Context, config *config.Config) (*redis.Redis, error) {
	r, err := redis.New(ctx, config.RedisConfig.GetRedisOptions())
	if err != nil {
		return nil, err
	}
	return r, nil
}

func InitializeZap(config *config.Config) (*zap.Zap, error) {
	z, err := zap.New(config.ZapConfig.GetZapConfig())
	if err != nil {
		return nil, err
	}
	return z, nil
}

func InitializeJwt(config *config.Config) (*jwt.Jwt, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	j, err := jwt.New(
		privateKey, config.JWTConfig.TokenDuration, config.JWTConfig.RefreshDuration, config.JWTConfig.RefreshBuffer,
	)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func InitializePrometheus(config *config.Config) *prometheus.Prometheus {
	return prometheus.New(
		config.PrometheusConfig.Namespace, config.PrometheusConfig.Subsystem, config.PrometheusConfig.MetricPath,
	)
}
