package modules

type BaseConfig struct {
	AppName    string `mapstructure:"app_name" yaml:"app_name" default:"data-collection-hub-server"`
	AppPort    string `mapstructure:"app_port" yaml:"app_port" default:"3000"`
	AppHost    string `mapstructure:"app_host" yaml:"app_host" default:"localhost"`
	AppVersion string `mapstructure:"app_version" yaml:"app_version" default:"1.0.0"`
}
