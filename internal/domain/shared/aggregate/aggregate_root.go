package aggregate

import "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"

type AggregateRoot struct {
	events []event.DomainEvent
}

func (a *AggregateRoot) AddEvent(e event.DomainEvent) {
	a.events = append(a.events, e)
}

func (a *AggregateRoot) PullEvents() []event.DomainEvent {

	res := a.events

	a.events = nil

	return res
}
