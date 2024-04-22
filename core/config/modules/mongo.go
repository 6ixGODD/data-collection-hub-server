package modules

type MongoConfig struct {
	Uri              string  `env:"MONGO_URI" mapstructure:"mongo_uri" yaml:"mongo_uri" default:"mongodb://localhost:27017"`
	Database         string  `env:"MONGO_DATABASE" mapstructure:"mongo_database" yaml:"mongo_database" default:"test"`
	ConnectTimeoutMS *int64  `env:"MONGO_CONNECT_TIMEOUT_MS" mapstructure:"mongo_connect_timeout_ms" yaml:"mongo_connect_timeout_ms" default:"10000"`
	MaxPoolSize      *uint64 `env:"MONGO_MAX_POOL_SIZE" mapstructure:"mongo_max_pool_size" yaml:"mongo_max_pool_size" default:"10"`
	MinPoolSize      *uint64 `env:"MONGO_MIN_POOL_SIZE" mapstructure:"mongo_min_pool_size" yaml:"mongo_min_pool_size" default:"1"`
	SocketTimeoutMS  *int64  `env:"MONGO_SOCKET_TIMEOUT_MS" mapstructure:"mongo_socket_timeout_ms" yaml:"mongo_socket_timeout_ms" default:"10000"`
}
