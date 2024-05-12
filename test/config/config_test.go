package config__test

import (
	"testing"

	"data-collection-hub-server/internal/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg, err := config.New()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	t.Logf("config: %+v", cfg)
}
