package allocation

import "errors"

var (
	ErrAllocationNotFound    = errors.New("allocation not found")
	ErrInvalidProductionUnit = errors.New("invalid production unit")
	ErrCycleAlreadyExists    = errors.New("active cycle already exists")
)
