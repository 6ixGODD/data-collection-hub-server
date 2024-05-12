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

func InitializeMongo(ctx context.Context) (*mongo.Mongo, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	m, err := mongo.New(ctx, cfg.MongoConfig.GetQmgoConfig(), cfg.MongoConfig.PingTimeoutS, cfg.MongoConfig.Database)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func InitializeRedis(ctx context.Context) (*redis.Redis, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	r, err := redis.New(ctx, cfg.RedisConfig.GetRedisOptions())
	if err != nil {
		return nil, err
	}
	return r, nil
}

func InitializeZap() (*zap.Zap, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	z, err := zap.New(cfg.ZapConfig.GetZapConfig())
	if err != nil {
		return nil, err
	}
	return z, nil
}

func InitializeJwt() (*jwt.Jwt, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	j, err := jwt.New(
		privateKey, cfg.JWTConfig.TokenDuration, cfg.JWTConfig.RefreshDuration, cfg.JWTConfig.RefreshBuffer,
	)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func InitializePrometheus() (*prometheus.Prometheus, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	p := prometheus.New(cfg.PrometheusConfig.Namespace, cfg.PrometheusConfig.Subsystem, cfg.PrometheusConfig.MetricPath)
	if err != nil {
		return nil, err
	}
	return p, nil
}
