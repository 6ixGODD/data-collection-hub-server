package modules

import (
	"time"
)

type LimiterConfig struct {
	Max        int           `env:"LIMITER_MAX" mapstructure:"limiter_max" yaml:"limiter_max" default:"20"`
	Expiration time.Duration `env:"LIMITER_EXPIRATION" mapstructure:"limiter_expiration" yaml:"limiter_expiration" default:"30s"`
}
