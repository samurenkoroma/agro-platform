package inmemory

import (
	"context"
	"sync"

	"github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	"github.com/samurenkoroma/agro-platform/internal/shared/bus"
)

type InMemoryEventBus struct {
	handlers map[string][]bus.EventHandler
	mu       sync.RWMutex
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make(map[string][]bus.EventHandler),
	}
}

func (b *InMemoryEventBus) Register(eventName string, handler bus.EventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *InMemoryEventBus) Publish(
	ctx context.Context,
	events []event.DomainEvent,
) error {

	for _, e := range events {

		b.mu.RLock()
		handlers := b.handlers[e.EventType()]
		b.mu.RUnlock()

		for _, h := range handlers {
			if err := h(ctx, e); err != nil {
				return err
			}
		}
	}

	return nil
}
