package conf_v3


import (
	"testing"
)


func TestLoadConfig(t *testing.T) {
	LoadConfig()

	t.Log(Cfg)
}