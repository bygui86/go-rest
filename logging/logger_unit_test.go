// +build unit

package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInitLogging(t *testing.T) {
	assert.NotNil(t, Config)
	assert.NotNil(t, Log)
	assert.NotNil(t, SugaredLog)

	assert.Equal(t, zap.NewAtomicLevelAt(defaultLevel), Config.Level)
	assert.Equal(t, defaultEncoding, Config.Encoding)

	assert.Equal(t, SugaredLog, Log.Sugar())
	assert.Equal(t, Log, SugaredLog.Desugar())

	assert.Equal(t, true, Log.Core().Enabled(defaultLevel))
	assert.Equal(t, true, Log.Core().Enabled(zapcore.WarnLevel))
	assert.Equal(t, true, Log.Core().Enabled(zapcore.ErrorLevel))
	assert.Equal(t, false, Log.Core().Enabled(zapcore.DebugLevel))
}
