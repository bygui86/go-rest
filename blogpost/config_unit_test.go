// +build unit

package blogpost

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_default(t *testing.T) {
	cfg := loadConfig()
	assert.Equal(t, restHostEnvVarDefault, cfg.RestHost)
	assert.Equal(t, restPortEnvVarDefault, cfg.RestPort)
}

func TestLoadConfig_custom(t *testing.T) {
	customHost := "custom-host"
	customPort := 8888

	os.Setenv(restHostEnvVar, customHost)
	os.Setenv(restPortEnvVar, strconv.Itoa(customPort))

	cfg := loadConfig()
	assert.Equal(t, customHost, cfg.RestHost)
	assert.Equal(t, customPort, cfg.RestPort)
}
