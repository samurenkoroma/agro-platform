package timeline

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventTimelineCreated   = "timeline.created"
	EventTimelineItemAdded = "timeline.item.added"
	EventTimelineArchived  = "timeline.archived"
)

type TimelineCreated struct {
	ev.BaseEvent
}

func NewTimelineCreated(id vo.ID) TimelineCreated {
	return TimelineCreated{
		BaseEvent: ev.NewBaseEvent(id, EventTimelineCreated),
	}
}

type ItemAdded struct {
	ev.BaseEvent
	ItemId vo.ID
}

func NewItemAdded(id vo.ID, itemId vo.ID) ItemAdded {
	return ItemAdded{
		BaseEvent: ev.NewBaseEvent(id, EventTimelineItemAdded),
		ItemId:    itemId,
	}
}

type TimelineArchived struct {
	ev.BaseEvent
}

func NewTimelineArchived(id vo.ID) TimelineArchived {
	return TimelineArchived{
		BaseEvent: ev.NewBaseEvent(id, EventTimelineArchived),
	}

}
