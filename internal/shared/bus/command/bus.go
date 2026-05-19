package command

import "context"

type Command interface {
	CommandName() string
}

type Bus interface {
	Dispatch(
		ctx context.Context,
		cmd Command,
	) error
}
