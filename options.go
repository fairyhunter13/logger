package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Options specifies the functional options to be used for the init of this package.
type Options func(inst *Logger)

// WithCore supply the Core to the underlying Logger.
func WithCore(core zapcore.Core) Options {
	return func(l *Logger) {
		l.Core = core
	}
}

// WithZapOptions supplies the slice of zap.Option to the underlying Logger.
func WithZapOptions(options []zap.Option) Options {
	return func(l *Logger) {
		l.Options = options
	}
}
