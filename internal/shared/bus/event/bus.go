package event

import "context"

type Event interface {
	EventName() string
}

type Bus interface {
	Publish(ctx context.Context, events ...Event) error
	Subscribe(name string, handler any)
}
