package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logEncodingDefault = "console"
	logLevelDefault    = "info"
)

var Log *zap.Logger
var SugaredLog *zap.SugaredLogger

func init() {
	level := zapcore.InfoLevel
	err := level.Set(logLevelDefault)
	if err != nil {
		panic(err)
	}
	buildLogger(logEncodingDefault, level)
}

func buildLogger(encoding string, level zapcore.Level) {
	Log, _ = zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			TimeKey:     "time",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}.Build()
	SugaredLog = Log.Sugar()
}
