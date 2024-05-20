package mods

import (
	"data-collection-hub-server/internal/pkg/config/mods/middleware"
)

type MiddlewareConfig struct {
	LimiterConfig middleware.LimiterConfig `mapstructure:"limiter" yaml:"limiter"`
	CorsConfig    middleware.CorsConfig    `mapstructure:"cors" yaml:"cors"`
	AuthConfig    middleware.AuthConfig    `mapstructure:"auth" yaml:"auth"`
}
