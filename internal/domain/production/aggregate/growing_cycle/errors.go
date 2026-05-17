package growingcycle

import "errors"

var (
	ErrInvalidTransition = errors.New("invalid status transition")

	ErrAlreadyStarted = errors.New("cycle already started")

	ErrAlreadyFinished = errors.New("cycle already finished")
)
