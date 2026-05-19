package startgrowingcycle

import "errors"

var (
	ErrProductionUnitNotFound = errors.New("production unit not found")
	ErrInvalidProductionUnit  = errors.New("invalid production unit")
	ErrCycleAlreadyExists     = errors.New("active cycle already exists")
)
