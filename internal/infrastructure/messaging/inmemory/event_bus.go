package inmemory

import (
	"context"
	"sync"

	"github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	"github.com/samurenkoroma/agro-platform/internal/shared/bus"
	"github.com/samurenkoroma/agro-platform/pkg/logger"
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
	log := logger.FromContext(ctx)
	for _, e := range events {

		b.mu.RLock()
		handlers := b.handlers[e.EventType()]
		b.mu.RUnlock()

		if len(handlers) == 0 {
			log.Debug("event_bus: no handlers", "event_type", e.EventType())
			continue
		}

		for _, h := range handlers {
			if err := h(ctx, e); err != nil {
				log.Error("event_bus: handler failed",
					"event_type", e.EventType(),
					"aggregate_id", e.AggregateID(),
					"error", err,
				)
				return err
			}
		}
		log.Debug("event_bus: published",
			"event_type", e.EventType(),
			"aggregate_id", e.AggregateID(),
			"handlers", len(handlers),
		)
	}

	return nil
}
