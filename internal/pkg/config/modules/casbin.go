package modules

type CasbinConfig struct {
	ModelPath        string `mapstructure:"casbin_model_path" yaml:"casbin_model_path" default:"./config/casbin_model.conf"`
	PolicyPath       string `mapstructure:"casbin_policy_path" yaml:"casbin_policy_path" default:"./config/casbin_policy.csv"`
	DbSpec           string `mapstructure:"casbin_db_spec" yaml:"casbin_db_spec" default:""`
	EnableLog        bool   `mapstructure:"casbin_enable_log" yaml:"casbin_enable_log" default:"false"`
	EnableHttp       bool   `mapstructure:"casbin_enable_http" yaml:"casbin_enable_http" default:"false"`
	HttpPort         int    `mapstructure:"casbin_http_port" yaml:"casbin_http_port" default:"80"`
	EnableAutoLoad   bool   `mapstructure:"casbin_enable_auto_load" yaml:"casbin_enable_auto_load" default:"false"`
	AutoLoadInternal int    `mapstructure:"casbin_auto_load_internal" yaml:"casbin_auto_load_internal" default:"60"`
}
