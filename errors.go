package logger

import "errors"

// List of all errors used in this package
var (
	ErrEmptyOpts = errors.New("Options must not be nil or empty")
)
