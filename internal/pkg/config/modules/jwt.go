package modules

type JWTConfig struct {
	SecretKey            string `mapstructure:"jwt_secret_key" yaml:"jwt_secret_key" default:"secret"`
	TokenExpireIn        int    `mapstructure:"jwt_token_expire_in" yaml:"jwt_token_expire_in" default:"3600"`
	RefreshTokenExpireIn int    `mapstructure:"jwt_refresh_token_expire_in" yaml:"jwt_refresh_token_expire_in" default:"7200"`
}
