package inmemory

import (
	"context"
	"fmt"

	cb "github.com/samurenkoroma/agro-platform/internal/shared/bus/command"
)

type CommandBus struct {
	registry *Registry
}

func NewCommandBus(reg *Registry) *CommandBus {
	return &CommandBus{
		registry: reg,
	}
}

func (b *CommandBus) Dispatch(ctx context.Context, cmd cb.Command) error {
	name := cmd.CommandName()

	raw, ok := b.registry.Get(name)

	if !ok {
		return cb.ErrHandlerNotFound
	}

	handler, ok := raw.(func(context.Context, cb.Command) error)

	if !ok {
		return fmt.Errorf(
			"invalid handler: %s",
			name,
		)
	}

	return handler(ctx, cmd)
}
