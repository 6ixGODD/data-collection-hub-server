package modules

type RedisConfig struct {
	Host     string `mapstructure:"host" env:"REDIS_HOST" default:"localhost"`
	Port     int    `mapstructure:"port" env:"REDIS_PORT" default:"6379"`
	Password string `mapstructure:"password" env:"REDIS_PASSWORD" default:""`
	DB       int    `mapstructure:"db" env:"REDIS_DB" default:"0"`
}
