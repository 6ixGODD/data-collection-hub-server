package config

import (
	"data-collection-hub-server/core/config/modules"
)

type Config struct {
	BaseConfig    *modules.BaseConfig
	CasbinConfig  *modules.CasbinConfig
	JWTConfig     *modules.JWTConfig
	MongoConfig   *modules.MongoConfig
	RedisConfig   *modules.RedisConfig
	ZapConfig     *modules.ZapConfig
	FiberConfig   *modules.FiberConfig
	LimiterConfig *modules.LimiterConfig
}
