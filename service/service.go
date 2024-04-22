package service

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Service is the interface for the service layer.
type Service struct {
	logger *zap.Logger
	config *viper.Viper
}
