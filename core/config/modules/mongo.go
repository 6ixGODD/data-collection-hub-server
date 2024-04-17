package modules

type MongoConfig struct {
	Uri      string `mapstructure:"uri" env:"MONGO_URI" default:"mongodb://localhost:27017"`
	Database string `mapstructure:"database" env:"MONGO_DATABASE" default:"test"`
	Username string `mapstructure:"username" env:"MONGO_USERNAME" default:""`
	Password string `mapstructure:"password" env:"MONGO_PASSWORD" default:""`
}
