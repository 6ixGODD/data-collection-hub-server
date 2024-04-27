package service

import (
	"data-collection-hub-server/internal/pkg/config"
	"go.uber.org/zap"
)

// Service is the interface for the service layer.
type Service struct {
	logger *zap.Logger
	config *config.Config
}

// NewService creates a new instance of Service.
func NewService(logger *zap.Logger, config *config.Config) *Service {
	return &Service{
		logger: logger,
		config: config,
	}
}
