package telemetry

import "errors"

var (
	ErrInvalidValue     = errors.New("invalid value")
	ErrInvalidTimestamp = errors.New("invalid timestamp")
	ErrMissingSensor    = errors.New("sensor required")
)
