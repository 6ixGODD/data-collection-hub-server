package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"data-collection-hub-server/core/config/modules"
)

var (
	logger *zap.Logger
)

func InitLogger(config *modules.ZapConfig) (err error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	logger, err = zap.Config{
		Level:             zap.NewAtomicLevelAt(transLevel(config.Level)),
		Development:       config.Development,
		Encoding:          config.Encoding,
		EncoderConfig:     encoderConfig,
		DisableStacktrace: config.DisableStacktrace,
		DisableCaller:     config.DisableCaller,
		OutputPaths:       config.OutputPaths,
		ErrorOutputPaths:  config.ErrorOutputPaths,
	}.Build()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func transLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func SyncLogger() { // flush log buffer
	if logger != nil {
		_ = logger.Sync()
	}
}

func GetLogger() *zap.Logger {
	return logger
}
