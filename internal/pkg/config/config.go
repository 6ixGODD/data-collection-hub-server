package config

import (
	"sync"

	"data-collection-hub-server/internal/pkg/config/mods"
	"github.com/mcuadros/go-defaults"
)

var (
	configInstance = &Config{}
	once           sync.Once
)

type Config struct {
	BaseConfig       mods.BaseConfig       `mapstructure:"base" yaml:"base"`
	CasbinConfig     mods.CasbinConfig     `mapstructure:"casbin" yaml:"casbin"`
	CorsConfig       mods.CorsConfig       `mapstructure:"cors" yaml:"cors"`
	FiberConfig      mods.FiberConfig      `mapstructure:"fiber" yaml:"fiber"`
	JWTConfig        mods.JWTConfig        `mapstructure:"jwt" yaml:"jwt"`
	LimiterConfig    mods.LimiterConfig    `mapstructure:"limiter" yaml:"limiter"`
	MongoConfig      mods.MongoConfig      `mapstructure:"mongo" yaml:"mongo"`
	PrometheusConfig mods.PrometheusConfig `mapstructure:"prometheus" yaml:"prometheus"`
	CacheConfig      mods.CacheConfig      `mapstructure:"cache" yaml:"cache"`
	RedisConfig      mods.RedisConfig      `mapstructure:"redis" yaml:"redis"`
	TasksConfig      mods.TasksConfig      `mapstructure:"tasks" yaml:"tasks"`
	ZapConfig        mods.ZapConfig        `mapstructure:"zap" yaml:"zap"`
}

// New returns instance of Config
func New() (config *Config, err error) {
	once.Do(
		func() {
			defaults.SetDefaults(configInstance)
		},
	)
	return configInstance, nil
}

func Update(cfg *Config) {
	configInstance = cfg
}
