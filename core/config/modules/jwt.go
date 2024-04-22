package modules

type JWTConfig struct {
	SecretKey            string `env:"JWT_SECRET_KEY" mapstructure:"jwt_secret_key" yaml:"jwt_secret_key" default:"secret"`
	TokenExpireIn        int    `env:"JWT_TOKEN_EXPIRE_IN" mapstructure:"jwt_token_expire_in" yaml:"jwt_token_expire_in" default:"3600"`
	RefreshTokenExpireIn int    `env:"JWT_REFRESH_TOKEN_EXPIRE_IN" mapstructure:"jwt_refresh_token_expire_in" yaml:"jwt_refresh_token_expire_in" default:"7200"`
}
