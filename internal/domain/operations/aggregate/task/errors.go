package task

import "errors"

var (
	ErrTaskNotFound     = errors.New("task not found")
	ErrInvalidStatus    = errors.New("invalid task status transition")
	ErrAlreadyCompleted = errors.New("task already completed")
	ErrAlreadyCancelled = errors.New("task already cancelled")
)
