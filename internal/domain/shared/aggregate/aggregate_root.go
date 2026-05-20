package aggregate

import "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"

type Aggregate interface {
	AddEvent(e event.DomainEvent)
	PullEvents() []event.DomainEvent
}

var _ Aggregate = (*BaseAggregate)(nil)

type BaseAggregate struct {
	events []event.DomainEvent
}

func (a *BaseAggregate) AddEvent(e event.DomainEvent) {
	a.events = append(a.events, e)
}

func (a *BaseAggregate) PullEvents() []event.DomainEvent {
	res := a.events
	a.events = nil
	return res
}
