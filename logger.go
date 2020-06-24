package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *Logger

// Logger specifies the logger instance for this package.
type Logger struct {
	logger  *zap.Logger
	Core    zapcore.Core
	Options []zap.Option
}

// Validate validate the logger parameters that is required.
func (l *Logger) Validate() (err error) {
	if l.Core == nil || len(l.Options) == 0 {
		err = ErrEmptyFuncOpts
	}
	return
}

// Init initialize the logger local instance of this package.
func Init(opts ...Options) (err error) {
	logger = new(Logger)
	logger.Core = zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), os.Stdout, zap.NewAtomicLevelAt(zapcore.DebugLevel))
	logger.Options = []zap.Option{zap.AddCaller()}

	for _, opt := range opts {
		err = opt(logger)
		if err != nil {
			return
		}
	}

	err = logger.Validate()
	if err != nil {
		return
	}

	logger.logger = zap.New(logger.Core, logger.Options...)

	return
}
