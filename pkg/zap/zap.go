package zap

import (
	"context"

	"data-collection-hub-server/internal/pkg/config/modules"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	MainTag      = "MAIN"
	MongoTag     = "MONGO"
	RequestTag   = "REQUEST"
	LoginTag     = "LOGIN"
	OperationTag = "OPERATION"
	SystemTag    = "SYSTEM"
)

type (
	tagKey       struct{}
	requestIDKey struct{}
	userIDKey    struct{}
	operationKey struct{}
)

func InitLogger(config *modules.ZapConfig) (err error) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	var level zapcore.Level
	if level, err = zapcore.ParseLevel(config.Level); err != nil {
		level = zapcore.InfoLevel
	}
	var baseLogger *zap.Logger
	baseLogger, err = zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      config.Development,
		Encoding:         config.Encoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      config.OutputPaths,
		ErrorOutputPaths: config.ErrorOutputPaths,
	}.Build()
	if err != nil {
		return err
	}

	baseLogger = baseLogger.WithOptions(
		zap.WithCaller(true),
		zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCallerSkip(config.CallerSkip),
	)
	zap.ReplaceGlobals(baseLogger)

	return nil
}

func SetTagInContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}

func getTagFromContext(ctx context.Context) string {
	if tag := ctx.Value(tagKey{}); tag != nil {
		if tagStr, ok := tag.(string); ok {
			return tagStr
		}
	}
	return MainTag
}

func SetRequestIDInContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

func getRequestIDFromContext(ctx context.Context) string {
	if requestID := ctx.Value(requestIDKey{}); requestID != nil {
		if requestID, ok := requestID.(string); ok {
			return requestID
		}
	}
	return ""
}

func SetUserIDWithContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func getUserIDFromContext(ctx context.Context) string {
	if userID := ctx.Value(userIDKey{}); userID != nil {
		if userID, ok := userID.(string); ok {
			return userID
		}
	}
	return ""
}

func SetOperationWithContext(ctx context.Context, operation string) context.Context {
	return context.WithValue(ctx, operationKey{}, operation)
}

func getOperationFromContext(ctx context.Context) string {
	if operation := ctx.Value(operationKey{}); operation != nil {
		if operation, ok := operation.(string); ok {
			return operation
		}
	}
	return ""
}

func GetLoggerWithContext(ctx context.Context) *zap.Logger {
	var Fields []zap.Field
	if tag := getTagFromContext(ctx); tag != "" {
		Fields = append(Fields, zap.String("tag", tag))
	}
	if reqID := getRequestIDFromContext(ctx); reqID != "" {
		Fields = append(Fields, zap.String("reqID", reqID))
	}
	if userID := getUserIDFromContext(ctx); userID != "" {
		Fields = append(Fields, zap.String("userID", userID))
	}
	if operation := getOperationFromContext(ctx); operation != "" {
		Fields = append(Fields, zap.String("operation", operation))
	}
	return zap.L().With(Fields...)
}
