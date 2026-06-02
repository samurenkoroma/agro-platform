package harvest

import "errors"

var (
	ErrPlantingNotFound      = errors.New("planting not found")
	ErrInvalidProductionUnit = errors.New("invalid production unit")
	ErrCycleAlreadyExists    = errors.New("active cycle already exists")
)
