package zap

import (
	"context"

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

	tag       = "TAG"
	requestID = "REQUEST_ID"
	userID    = "USER_ID"
	operation = "OPERATION"
)

type Logger struct {
	Logger  *zap.Logger
	Config  *zap.Config
	Options []zap.Option
}

func New(config *zap.Config, options ...zap.Option) (l *Logger, err error) {
	l = &Logger{
		Config:  config,
		Options: options,
	}
	if err := l.Init(); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *Logger) Init() error {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	l.Config.EncoderConfig = encoderConfig
	logger, err := l.Config.Build()
	if err != nil {
		return err
	}

	l.Options = append(l.Options, zap.WithCaller(true), zap.AddStacktrace(zap.ErrorLevel))
	l.Logger = logger.WithOptions(l.Options...)

	return nil
}

func (l *Logger) GetLogger(ctx context.Context) (logger *zap.Logger, err error) {
	var fields []zap.Field
	if tag := l.getTagFromContext(ctx); tag != "" {
		fields = append(fields, zap.String("tag", tag))
	}
	if reqID := l.getRequestIDFromContext(ctx); reqID != "" {
		fields = append(fields, zap.String("reqID", reqID))
	}
	if userID := l.getUserIDFromContext(ctx); userID != "" {
		fields = append(fields, zap.String("userID", userID))
	}
	if operation := l.getOperationFromContext(ctx); operation != "" {
		fields = append(fields, zap.String("operation", operation))
	}
	if l.Logger == nil {
		if err := l.Init(); err != nil {
			return nil, err
		}
	}
	return l.Logger.With(fields...), nil
}

func (l *Logger) SetTagInContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tag, tag)
}

func (l *Logger) getTagFromContext(ctx context.Context) string {
	if tag := ctx.Value(tag); tag != nil {
		if tagStr, ok := tag.(string); ok {
			return tagStr
		}
	}
	return MainTag
}

func (l *Logger) SetRequestIDInContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestID, requestID)
}

func (l *Logger) getRequestIDFromContext(ctx context.Context) string {
	if requestID := ctx.Value(requestID); requestID != nil {
		if requestID, ok := requestID.(string); ok {
			return requestID
		}
	}
	return ""
}

func (l *Logger) SetUserIDWithContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userID, userID)
}

func (l *Logger) getUserIDFromContext(ctx context.Context) string {
	if userID := ctx.Value(userID); userID != nil {
		if userID, ok := userID.(string); ok {
			return userID
		}
	}
	return ""
}

func (l *Logger) SetOperationWithContext(ctx context.Context, operation string) context.Context {
	return context.WithValue(ctx, operation, operation)
}

func (l *Logger) getOperationFromContext(ctx context.Context) string {
	if operation := ctx.Value(operation); operation != nil {
		if operation, ok := operation.(string); ok {
			return operation
		}
	}
	return ""
}
