package growingcycle

import "errors"

var (
	ErrGrowingCycleNotFound  = errors.New("growing cycle not found")
	ErrInvalidProductionUnit = errors.New("invalid production unit")
	ErrCycleAlreadyExists    = errors.New("active cycle already exists")
)
