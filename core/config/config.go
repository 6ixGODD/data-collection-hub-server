package config

import (
	"data-collection-hub-server/core/config/modules"
)

type Config struct {
	CORSConfig      *modules.CORSConfig
	JWTConfig       *modules.JWTConfig
	MemcachedConfig *modules.MemcachedConfig
	MongoConfig     *modules.MongoConfig
	RedisConfig     *modules.RedisConfig
	ZapConfig       *modules.ZapConfig
}

func GetConfig() *Config {
	return &Config{
		CORSConfig:      &modules.CORSConfig{},
		JWTConfig:       &modules.JWTConfig{},
		MemcachedConfig: &modules.MemcachedConfig{},
		MongoConfig:     &modules.MongoConfig{},
		RedisConfig:     &modules.RedisConfig{},
		ZapConfig:       &modules.ZapConfig{},
	}
}
