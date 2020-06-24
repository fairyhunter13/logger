package logger

import "errors"

// List of all errors used in this package
var (
	ErrEmptyFuncOpts = errors.New("Functional options must not be nil or empty")
)
