package event

type AggregateRoot struct {
	events []DomainEvent
}

func (a *AggregateRoot) AddEvent(
	e DomainEvent,
) {
	a.events = append(
		a.events,
		e,
	)
}

func (a *AggregateRoot) PullEvents() []DomainEvent {

	res := a.events

	a.events = nil

	return res
}
