package command

import "errors"

var (
	ErrHandlerNotFound = errors.New("command handler not found")

	ErrInvalidCommand = errors.New("invalid command")
)
