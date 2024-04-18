package config

import (
	"data-collection-hub-server/core/config/modules"
)

type Config struct {
	CasbinConfig    *modules.CasbinConfig
	JWTConfig       *modules.JWTConfig
	MemcachedConfig *modules.MemcachedConfig
	MongoConfig     *modules.MongoConfig
	RedisConfig     *modules.RedisConfig
	ZapConfig       *modules.ZapConfig
}

func GetConfig() *Config { // Viper read ENV
	return &Config{
		CasbinConfig:    &modules.CasbinConfig{},
		JWTConfig:       &modules.JWTConfig{},
		MemcachedConfig: &modules.MemcachedConfig{},
		MongoConfig:     &modules.MongoConfig{},
		RedisConfig:     &modules.RedisConfig{},
		ZapConfig:       &modules.ZapConfig{},
	}
}
