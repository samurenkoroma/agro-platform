package bus

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
)

type EventHandler func(ctx context.Context, e event.DomainEvent) error

type EventBus interface {
	Register(eventName string, handler EventHandler)
	Publish(context.Context, []event.DomainEvent) error
}
