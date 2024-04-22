package modules

import (
	"time"
)

type FiberConfig struct {
	Prefork                 bool          `env:"FIBER_PREFORK" mapstructure:"prefork" yaml:"prefork" default:"false"`
	ServerHeader            string        `env:"FIBER_SERVER_HEADER" mapstructure:"server_header" yaml:"server_header" default:""`
	BodyLimit               int           `env:"FIBER_BODY_LIMIT" mapstructure:"body_limit" yaml:"body_limit" default:"4 * 1024 * 1024"`
	Concurrency             int           `env:"FIBER_CONCURRENCY" mapstructure:"concurrency" yaml:"concurrency" default:"256 * 1024"`
	ReadTimeout             time.Duration `env:"FIBER_READ_TIMEOUT" mapstructure:"read_timeout" yaml:"read_timeout" default:"10s"`
	WriteTimeout            time.Duration `env:"FIBER_WRITE_TIMEOUT" mapstructure:"write_timeout" yaml:"write_timeout" default:"10s"`
	IdleTimeout             time.Duration `env:"FIBER_IDLE_TIMEOUT" mapstructure:"idle_timeout" yaml:"idle_timeout" default:"2m"`
	ReadBufferSize          int           `env:"FIBER_READ_BUFFER_SIZE" mapstructure:"read_buffer_size" yaml:"read_buffer_size" default:"4096"`
	WriteBufferSize         int           `env:"FIBER_WRITE_BUFFER_SIZE" mapstructure:"write_buffer_size" yaml:"write_buffer_size" default:"4096"`
	ProxyHeader             string        `env:"FIBER_PROXY_HEADER" mapstructure:"proxy_header" yaml:"proxy_header" default:"X-Forwarded-For"`
	DisableKeepalive        bool          `env:"FIBER_DISABLE_KEEPALIVE" mapstructure:"disable_keepalive" yaml:"disable_keepalive" default:"false"`
	DisableStartupMessage   bool          `env:"FIBER_DISABLE_STARTUP_MESSAGE" mapstructure:"disable_startup_message" yaml:"disable_startup_message" default:"true"`
	ReduceMemoryUsage       bool          `env:"FIBER_REDUCE_MEMORY_USAGE" mapstructure:"reduce_memory_usage" yaml:"reduce_memory_usage" default:"false"`
	EnableTrustedProxyCheck bool          `env:"FIBER_ENABLE_TRUSTED_PROXY_CHECK" mapstructure:"enable_trusted_proxy_check" yaml:"enable_trusted_proxy_check" default:"false"`
	TrustedProxies          []string      `env:"FIBER_TRUSTED_PROXIES" mapstructure:"trusted_proxies" yaml:"trusted_proxies" default:""`
	EnablePrintRoutes       bool          `env:"FIBER_ENABLE_PRINT_ROUTES" mapstructure:"enable_print_routes" yaml:"enable_print_routes" default:"true"`
}
