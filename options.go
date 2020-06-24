package logger

// Options specifies the functional options to be used for the init of this package.
type Options func(inst *Logger) (err error)
