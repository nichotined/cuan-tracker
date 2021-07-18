package config

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	var (
		_, b, _, _ = runtime.Caller(0)
		Root       = filepath.Join(filepath.Dir(b), "../../..")
	)
	cfg := GetConfig(Root + "/config.yml")
	assert.NotNil(t, cfg.Server.Port)
}
