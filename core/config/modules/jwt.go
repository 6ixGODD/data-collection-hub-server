package modules

type JWTConfig struct {
	SecretKey            string `mapstructure:"secret_key" env:"JWT_SECRET_KEY" default:"secret"`
	TokenExpireIn        int    `mapstructure:"token_expire_in" env:"JWT_TOKEN_EXPIRE_IN" default:"3600"`
	RefreshTokenExpireIn int    `mapstructure:"refresh_token" env:"JWT_REFRESH_TOKEN_EXPIRE_IN" default:"7200"`
}
