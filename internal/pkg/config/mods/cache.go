package mods

import (
	"time"
)

type CacheConfig struct {
	DefaultTTL            time.Duration `mapstructure:"default_ttl" yaml:"default_ttl" default:"5m"`
	LogCacheTTL           time.Duration `mapstructure:"log_cache_ttl" yaml:"log_cache_ttl" default:"5m"`
	UserCacheTTL          time.Duration `mapstructure:"user_cache_ttl" yaml:"user_cache_ttl" default:"5m"`
	NoticeCacheTTL        time.Duration `mapstructure:"notice_cache_ttl" yaml:"notice_cache_ttl" default:"5m"`
	DocumentationCacheTTL time.Duration `mapstructure:"documentation_cache_ttl" yaml:"documentation_cache_ttl" default:"5m"`
	RedisConfig           RedisConfig   `mapstructure:"redis" yaml:"redis"`
}
