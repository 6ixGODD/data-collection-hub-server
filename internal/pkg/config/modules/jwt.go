package modules

import (
	"time"
)

type JWTConfig struct {
	TokenDuration   time.Duration `mapstructure:"jwt_token_duration" yaml:"jwt_token_duration" default:"7200s"`
	RefreshDuration time.Duration `mapstructure:"jwt_refresh_duration" yaml:"jwt_refresh_duration" default:"14400s"`
}
