package modules

type CORSConfig struct {
	AllowOrigins     []string `mapstructure:"allow-origins" env:"CORS_ALLOW_ORIGINS" default:"*"`
	AllowMethods     []string `mapstructure:"allow-methods" env:"CORS_ALLOW_METHODS" default:"GET,POST,PUT,DELETE,PATCH,OPTIONS"`
	AllowHeaders     []string `mapstructure:"allow-headers" env:"CORS_ALLOW_HEADERS" default:"Origin,Content-Length,Content-Type,Authorization"`
	ExposeHeaders    []string `mapstructure:"expose-headers" env:"CORS_EXPOSE_HEADERS" default:""`
	AllowCredentials bool     `mapstructure:"allow-credentials" env:"CORS_ALLOW_CREDENTIALS" default:"true"`
	MaxAge           int      `mapstructure:"max-age" env:"CORS_MAX_AGE" default:"12"`
}
