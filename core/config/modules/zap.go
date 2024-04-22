package modules

type ZapConfig struct {
	Level             string   `env:"ZAP_LEVEL" mapstructure:"zap_level" yaml:"zap_level" default:"info"`
	Encoding          string   `env:"ZAP_ENCODING" mapstructure:"zap_encoding" yaml:"zap_encoding" default:"mapstructure"`
	EncoderLevel      string   `env:"ZAP_ENCODER_LEVEL" mapstructure:"zap_encoder_level" yaml:"zap_encoder_level" default:"capital"`
	Development       bool     `env:"ZAP_DEVELOPMENT" mapstructure:"zap_development" yaml:"zap_development" default:"false"`
	DisableCaller     bool     `env:"ZAP_DISABLE_CALLER" mapstructure:"zap_disable_caller" yaml:"zap_disable_caller" default:"false"`
	DisableStacktrace bool     `env:"ZAP_DISABLE_STACKTRACE" mapstructure:"zap_disable_stacktrace" yaml:"zap_disable_stacktrace" default:"false"`
	OutputPaths       []string `env:"ZAP_OUTPUT_PATHS" mapstructure:"zap_output_paths" yaml:"zap_output_paths" default:"stdout"`
	ErrorOutputPaths  []string `env:"ZAP_ERROR_OUTPUT_PATHS" mapstructure:"zap_error_output_paths" yaml:"zap_error_output_paths" default:"stderr"`
}
