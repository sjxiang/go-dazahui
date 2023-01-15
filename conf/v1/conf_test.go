package conf_v1

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config := LoadConfig()
	t.Log(*config)
}