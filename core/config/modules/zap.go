package modules

type ZapConfig struct {
	Level             string `mapstructure:"level" env:"ZAP_LEVEL" default:"info"`
	Format            string `mapstructure:"format" env:"ZAP_FORMAT" default:"json"`
	DisableCaller     bool   `mapstructure:"disable_caller" env:"ZAP_DISABLE_CALLER" default:"false"`
	DisableStacktrace bool   `mapstructure:"disable_stacktrace" env:"ZAP_DISABLE_STACKTRACE" default:"false"`
	Development       bool   `mapstructure:"development" env:"ZAP_DEVELOPMENT" default:"false"`
	DisableSampling   bool   `mapstructure:"disable_sampling" env:"ZAP_DISABLE_SAMPLING" default:"false"`
}
