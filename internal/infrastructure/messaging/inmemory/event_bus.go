package inmemory

import (
	"context"
	"sync"

	eb "github.com/samurenkoroma/agro-platform/internal/shared/bus/event"
)

type EventBus struct {
	mu sync.RWMutex

	handlers map[string][]func(context.Context, eb.Event) error
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]func(context.Context, eb.Event) error),
	}
}

func (b *EventBus) Subscribe(name string, handler any) {
	h := handler.(func(context.Context, eb.Event) error)

	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[name] = append(b.handlers[name], h)
}

func (b *EventBus) Publish(ctx context.Context, events ...eb.Event) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, event := range events {
		name := event.EventName()

		handlers := b.handlers[name]

		for _, h := range handlers {
			err := h(ctx, event)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
