package mods

import (
	"time"
)

type IdempotencyConfig struct {
	HeaderKey string        `mapstructure:"idempotency_header_key" yaml:"idempotency_header_key" default:"Idempotency-Key"`
	Expiry    time.Duration `mapstructure:"idempotency_expiry" yaml:"idempotency_expiry" default:"5m"`
}
