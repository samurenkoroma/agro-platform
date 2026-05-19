package command

import "context"

type Middleware func(ctx context.Context, cmd Command, next Next) error

type Next func(ctx context.Context, cmd Command) error
