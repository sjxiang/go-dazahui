package conf_v2

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	Init()
	
	assert.Equal(t, "127.0.0.1", os.Getenv("HOST"), "预期结果为 127.0.0.1 ")
}