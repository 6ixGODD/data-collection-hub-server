package viper__test

import (
	"encoding/json"
	"testing"

	"data-collection-hub-server/internal/pkg/config"
)

func TestViper(t *testing.T) {
	t.Log("TestViper")

	cfg, _ := config.New()
	t.Log(cfg)

	cfgBytes, _ := json.Marshal(cfg)
	t.Log(string(cfgBytes))

}
