//go:build test
// +build test

package zap_test

import (
	"context"
	"testing"

	z "data-collection-hub-server/pkg/zap"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestZap(t *testing.T) {
	ctx := context.Background()
	zapConfig := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zp, err := z.New(zapConfig)
	assert.NoError(t, err)
	ctx = zp.SetTagInContext(ctx, "TestZap__New")
	ctx = zp.SetRequestIDInContext(ctx, "TestZap__New_RequestID")
	ctx = zp.SetUserIDInContext(ctx, "TestZap__New_UserID")
	logger, err := zp.GetLogger(ctx)
	assert.NoError(t, err)
	logger.Info("TestZap__New: Info log")
	logger.Warn("TestZap__New: Warn log")
	logger.Error("TestZap__New: Error log")
}
