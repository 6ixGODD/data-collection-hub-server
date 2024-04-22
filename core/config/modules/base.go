package modules

type BaseConfig struct {
	AppName    string `env:"APP_NAME" mapstructure:"app_name" yaml:"app_name" default:"data-collection-hub-server"`
	AppEnv     string `env:"APP_ENV" mapstructure:"app_env" yaml:"app_env" default:"development"`
	AppPort    string `env:"APP_PORT" mapstructure:"app_port" yaml:"app_port" default:":3000"`
	AppHost    string `env:"APP_HOST" mapstructure:"app_host" yaml:"app_host" default:"localhost"`
	AppVersion string `env:"APP_VERSION" mapstructure:"app_version" yaml:"app_version" default:"v1"`
}
