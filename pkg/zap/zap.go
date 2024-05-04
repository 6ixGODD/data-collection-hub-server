package zap

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	MainTag      = "MAIN"
	SystemTag    = "SYSTEM"
	MongoTag     = "MONGO"
	RequestTag   = "REQUEST"
	LoginTag     = "LOGIN"
	OperationTag = "OPERATION"

	tag       = "TAG"
	requestID = "REQUEST_ID"
	userID    = "USER_ID"
	operation = "OPERATION"
)

var zapInstance *Zap

type Zap struct {
	Logger  *zap.Logger
	Config  *zap.Config
	Options []zap.Option
}

func New(config *zap.Config, options ...zap.Option) (z *Zap, err error) {
	if zapInstance != nil {
		return zapInstance, nil
	}
	zapInstance = &Zap{
		Config:  config,
		Options: options,
	}
	if err := zapInstance.Init(); err != nil {
		return nil, err
	}
	return zapInstance, nil
}

func (z *Zap) Init() error {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	z.Config.EncoderConfig = encoderConfig
	logger, err := z.Config.Build()
	if err != nil {
		return err
	}

	z.Options = append(z.Options, zap.WithCaller(true), zap.AddStacktrace(zap.ErrorLevel))
	z.Logger = logger.WithOptions(z.Options...)

	return nil
}

func (z *Zap) GetLogger(ctx context.Context) (logger *zap.Logger, err error) {
	var fields []zap.Field
	if tag := z.getTagFromContext(ctx); tag != "" {
		fields = append(fields, zap.String("tag", tag))
	}
	if reqID := z.getRequestIDFromContext(ctx); reqID != "" {
		fields = append(fields, zap.String("reqID", reqID))
	}
	if userID := z.getUserIDFromContext(ctx); userID != "" {
		fields = append(fields, zap.String("userID", userID))
	}
	if operation := z.getOperationFromContext(ctx); operation != "" {
		fields = append(fields, zap.String("operation", operation))
	}
	if z.Logger == nil {
		if err := z.Init(); err != nil {
			return nil, err
		}
	}
	return z.Logger.With(fields...), nil
}

func (z *Zap) SetTagInContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tag, tag)
}

func (z *Zap) getTagFromContext(ctx context.Context) string {
	if tag := ctx.Value(tag); tag != nil {
		if tagStr, ok := tag.(string); ok {
			return tagStr
		}
	}
	return MainTag
}

func (z *Zap) SetRequestIDInContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestID, requestID)
}

func (z *Zap) getRequestIDFromContext(ctx context.Context) string {
	if requestID := ctx.Value(requestID); requestID != nil {
		if requestID, ok := requestID.(string); ok {
			return requestID
		}
	}
	return ""
}

func (z *Zap) SetUserIDWithContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userID, userID)
}

func (z *Zap) getUserIDFromContext(ctx context.Context) string {
	if userID := ctx.Value(userID); userID != nil {
		if userID, ok := userID.(string); ok {
			return userID
		}
	}
	return ""
}

func (z *Zap) SetOperationWithContext(ctx context.Context, operation string) context.Context {
	return context.WithValue(ctx, operation, operation)
}

func (z *Zap) getOperationFromContext(ctx context.Context) string {
	if operation := ctx.Value(operation); operation != nil {
		if operation, ok := operation.(string); ok {
			return operation
		}
	}
	return ""
}
