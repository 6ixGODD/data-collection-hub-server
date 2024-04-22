package modules

import (
	"time"
)

type RedisConfig struct {
	Addr            string        `env:"REDIS_ADDR" mapstructure:"redis_addr" yaml:"redis_addr" default:"localhost:6379"`
	ClientName      string        `env:"REDIS_CLIENT_NAME" mapstructure:"redis_client_name" yaml:"redis_client_name" default:""`
	Username        string        `env:"REDIS_USERNAME" mapstructure:"redis_username" yaml:"redis_username" default:""`
	Password        string        `env:"REDIS_PASSWORD" mapstructure:"redis_password" yaml:"redis_password" default:""`
	DB              int           `env:"REDIS_DB" mapstructure:"redis_db" yaml:"redis_db" default:"0"`
	MaxRetries      int           `env:"REDIS_MAX_RETRIES" mapstructure:"redis_max_retries" yaml:"redis_max_retries" default:"5"`
	MinRetryBackoff time.Duration `env:"REDIS_MIN_RETRY_BACKOFF" mapstructure:"redis_min_retry_backoff" yaml:"redis_min_retry_backoff" default:"8ms"`
	MaxRetryBackoff time.Duration `env:"REDIS_MAX_RETRY_BACKOFF" mapstructure:"redis_max_retry_backoff" yaml:"redis_max_retry_backoff" default:"512ms"`
	DialTimeout     time.Duration `env:"REDIS_DIAL_TIMEOUT" mapstructure:"redis_dial_timeout" yaml:"redis_dial_timeout" default:"5s"`
	ReadTimeout     time.Duration `env:"REDIS_READ_TIMEOUT" mapstructure:"redis_read_timeout" yaml:"redis_read_timeout" default:"3s"`
	WriteTimeout    time.Duration `env:"REDIS_WRITE_TIMEOUT" mapstructure:"redis_write_timeout" yaml:"redis_write_timeout" default:"3s"`
	PoolSize        int           `env:"REDIS_POOL_SIZE" mapstructure:"redis_pool_size" yaml:"redis_pool_size" default:"10"`
	PoolTimeout     time.Duration `env:"REDIS_POOL_TIMEOUT" mapstructure:"redis_pool_timeout" yaml:"redis_pool_timeout" default:"4s"`
	MinIdleConns    int           `env:"REDIS_MIN_IDLE_CONNS" mapstructure:"redis_min_idle_conns" yaml:"redis_min_idle_conns" default:"0"`
	MaxIdleConns    int           `env:"REDIS_MAX_IDLE_CONNS" mapstructure:"redis_max_idle_conns" yaml:"redis_max_idle_conns" default:"0"`
	MaxActiveConns  int           `env:"REDIS_MAX_ACTIVE_CONNS" mapstructure:"redis_max_active_conns" yaml:"redis_max_active_conns" default:"30m"`
	ConnMaxIdleTime time.Duration `env:"REDIS_CONN_MAX_IDLE_TIME" mapstructure:"redis_conn_max_idle_time" yaml:"redis_conn_max_idle_time" default:"30m"`
	ConnMaxLifetime time.Duration `env:"REDIS_CONN_MAX_LIFETIME" mapstructure:"redis_conn_max_lifetime" yaml:"redis_conn_max_lifetime" default:"-1"`
}
