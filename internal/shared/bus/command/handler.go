package command

import "context"

type Handler[T Command] interface {
	Handle(ctx context.Context, cmd T) error
}
