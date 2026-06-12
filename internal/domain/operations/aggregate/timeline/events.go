package timeline

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventTimelineCreated   = "timeline.created"
	EventTimelineItemAdded = "timeline.item_added"
	EventTimelineArchived  = "timeline.archived"
)

type TimelineCreated struct{ ev.BaseEvent }

func NewTimelineCreated(id vo.ID) TimelineCreated {
	return TimelineCreated{ev.NewBaseEvent(id, EventTimelineCreated)}
}

type TimelineItemAdded struct {
	ev.BaseEvent
	ItemID vo.ID
}

func NewItemAdded(id, itemID vo.ID) TimelineItemAdded {
	return TimelineItemAdded{BaseEvent: ev.NewBaseEvent(id, EventTimelineItemAdded), ItemID: itemID}
}

type TimelineArchived struct{ ev.BaseEvent }

func NewTimelineArchived(id vo.ID) TimelineArchived {
	return TimelineArchived{ev.NewBaseEvent(id, EventTimelineArchived)}
}
