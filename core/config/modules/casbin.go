package modules

type CasbinConfig struct {
	ModelPath        string `env:"CASBIN_MODEL_PATH" mapstructure:"casbin_model_path" yaml:"casbin_model_path" default:"./config/casbin_model.conf"`
	PolicyPath       string `env:"CASBIN_POLICY_PATH" mapstructure:"casbin_policy_path" yaml:"casbin_policy_path" default:"./config/casbin_policy.csv"`
	DbSpec           string `env:"CASBIN_DB_SPEC" mapstructure:"casbin_db_spec" yaml:"casbin_db_spec" default:""`
	EnableLog        bool   `env:"CASBIN_ENABLE_LOG" mapstructure:"casbin_enable_log" yaml:"casbin_enable_log" default:"false"`
	EnableHttp       bool   `env:"CASBIN_ENABLE_HTTP" mapstructure:"casbin_enable_http" yaml:"casbin_enable_http" default:"false"`
	HttpPort         int    `env:"CASBIN_HTTP_PORT" mapstructure:"casbin_http_port" yaml:"casbin_http_port" default:"80"`
	EnableAutoLoad   bool   `env:"CASBIN_ENABLE_AUTO_LOAD" mapstructure:"casbin_enable_auto_load" yaml:"casbin_enable_auto_load" default:"false"`
	AutoLoadInternal int    `env:"CASBIN_AUTO_LOAD_INTERNAL" mapstructure:"casbin_auto_load_internal" yaml:"casbin_auto_load_internal" default:"60"`
}
