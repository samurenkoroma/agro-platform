package slot

import "errors"

var (
	ErrSlotFull = errors.New("slot full")

	ErrSlotBlocked = errors.New("slot blocked")

	ErrInvalidCount = errors.New("invalid plant count")
)
