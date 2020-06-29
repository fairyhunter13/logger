package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *Logger

// Logger specifies the logger instance for this package.
type Logger struct {
	*zap.Logger
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
	logger.Options = []zap.Option{zap.AddCaller(), zap.AddCallerSkip(1)}

	for _, opt := range opts {
		opt(logger)
	}

	err = logger.Validate()
	if err != nil {
		return
	}

	logger.Logger = zap.New(logger.Core, logger.Options...)

	return
}

func initLogger() {
	if logger == nil {
		Init()
	}
}

// Get returns the underlying logger
func Get() *Logger {
	initLogger()
	return logger
}

// Check returns a CheckedEntry if logging a message at the specified level is enabled.
func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	initLogger()
	return logger.Check(lvl, msg)
}

// DPanic logs message at DPanic level.
// If logger is in the development mode, the logger will be panic.
func DPanic(msg string, fields ...zap.Field) {
	initLogger()
	logger.DPanic(msg, fields...)
}

// Debug logs message at Debug level.
func Debug(msg string, fields ...zap.Field) {
	initLogger()
	logger.Debug(msg, fields...)
}

// Error logs message at Error level.
func Error(msg string, fields ...zap.Field) {
	initLogger()
	logger.Error(msg, fields...)
}

// Fatal logs message at Fatal level.
func Fatal(msg string, fields ...zap.Field) {
	initLogger()
	logger.Fatal(msg, fields...)
}

// Info logs message at Info level.
func Info(msg string, fields ...zap.Field) {
	initLogger()
	logger.Info(msg, fields...)
}

// Named adds a new path segment to the logger's name.
func Named(s string) *zap.Logger {
	initLogger()
	return logger.Named(s)
}

// Panic logs message at Panic level.
func Panic(msg string, fields ...zap.Field) {
	initLogger()
	logger.Panic(msg, fields...)
}

// Sugar returns the sugared logger.
func Sugar() *zap.SugaredLogger {
	initLogger()
	return logger.Sugar()
}

// Sync flushes the buffer/resources inside the logger.
func Sync() error {
	initLogger()
	return logger.Sync()
}

// Warn logs message at Warn level.
func Warn(msg string, fields ...zap.Field) {
	initLogger()
	logger.Warn(msg, fields...)
}

// With creates a child logger and adds structured context to it.
func With(fields ...zap.Field) *zap.Logger {
	initLogger()
	return logger.With(fields...)
}

// WithOptions clones the current logger and supplies it with options.
func WithOptions(options ...zap.Option) *zap.Logger {
	initLogger()
	return logger.WithOptions(options...)
}
