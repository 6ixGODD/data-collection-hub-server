package modules

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapConfig struct {
	Level            string   `mapstructure:"zap_level" yaml:"zap_level" default:"info"`
	Encoding         string   `mapstructure:"zap_encoding" yaml:"zap_encoding" default:"console"`
	Development      bool     `mapstructure:"zap_development" yaml:"zap_development" default:"false"`
	OutputPaths      []string `mapstructure:"zap_output_paths" yaml:"zap_output_paths" default:"stdout"`
	ErrorOutputPaths []string `mapstructure:"zap_error_output_paths" yaml:"zap_error_output_paths" default:"stderr"`
	CallerSkip       int      `mapstructure:"zap_caller_skip" yaml:"zap_caller_skip" default:"1"`
}

func (c *ZapConfig) ToZapConfig() *zap.Config {
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      c.Development,
		Encoding:         c.Encoding,
		OutputPaths:      c.OutputPaths,
		ErrorOutputPaths: c.ErrorOutputPaths,
	}
}
