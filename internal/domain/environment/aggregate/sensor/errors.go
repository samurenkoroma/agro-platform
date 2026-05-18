package sensor

import "errors"

var (
	ErrNilTimestamp    = errors.New("timestamp required")
	ErrArchivedSensor  = errors.New("sensor archived")
	ErrInvalidStatus   = errors.New("invalid sensor status")
	ErrAlreadyAttached = errors.New("sensor already attached")
)
