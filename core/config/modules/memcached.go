package modules

type MemcachedConfig struct {
	Host     string `mapstructure:"host" env:"MEMCACHED_HOST" default:"localhost"`
	Port     int    `mapstructure:"port" env:"MEMCACHED_PORT" default:"11211"`
	Username string `mapstructure:"username" env:"MEMCACHED_USERNAME" default:""`
	Password string `mapstructure:"password" env:"MEMCACHED_PASSWORD" default:""`
}
