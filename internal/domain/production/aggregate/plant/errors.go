package plant

import "errors"

var (
	ErrInvalidTransition = errors.New("invalid transition")

	ErrAlreadyHarvested = errors.New("already harvested")

	ErrAlreadyDiscarded = errors.New("already discarded")

	ErrPlantDead = errors.New("plant dead")
)
