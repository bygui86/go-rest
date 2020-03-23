// +build unit

package utils

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	stringEnvVarKey      = "STRING_TEST"
	stringEnvVarValue    = "sample-value"
	stringEnvVarFallback = "fallback-value"

	intEnvVarKey      = "INT_TEST"
	intEnvVarValue    = 86
	intEnvVarFallback = 42
)

func TestGetStringEnv(t *testing.T) {
	os.Setenv(stringEnvVarKey, stringEnvVarValue)

	value := GetStringEnv(stringEnvVarKey, stringEnvVarFallback)

	assert.Equal(t, stringEnvVarValue, value)

	// cleanup
	os.Unsetenv(stringEnvVarKey)
}

func TestGetStringEnv_fallback(t *testing.T) {
	value := GetStringEnv(stringEnvVarKey, stringEnvVarFallback)

	assert.Equal(t, stringEnvVarFallback, value)
}

func TestGetIntEnv(t *testing.T) {
	os.Setenv(intEnvVarKey, strconv.Itoa(intEnvVarValue))

	value := GetIntEnv(intEnvVarKey, intEnvVarFallback)

	assert.Equal(t, intEnvVarValue, value)

	// cleanup
	os.Unsetenv(intEnvVarKey)
}

func TestGetIntEnv_fallback(t *testing.T) {
	value := GetIntEnv(intEnvVarKey, intEnvVarFallback)

	assert.Equal(t, intEnvVarFallback, value)
}
