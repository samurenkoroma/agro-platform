package variety

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventVarietyCreated   = "variety.created"
	EventMaturityUpdated  = "variety.maturity.updated"
	EventGrowthUpdated    = "variety.growth.updated"
	EventSpacingUpdated   = "variety.spacing.updated"
	EventHarvestUpdated   = "variety.harvest.updated"
	EventYieldUpdated     = "variety.yield.updated"
	EventToleranceUpdated = "variety.tolerance.updated"
	EventVarietyArchived  = "variety.archived"
)

type VarietyCreated struct {
	ev.BaseEvent
}

type MaturityUpdated struct {
	ev.BaseEvent
}

type GrowthUpdated struct {
	ev.BaseEvent
}
type SpacingUpdated struct {
	ev.BaseEvent
}
type HarvestUpdated struct {
	ev.BaseEvent
}
type YieldUpdated struct {
	ev.BaseEvent
}
type ToleranceUpdated struct {
	ev.BaseEvent
}
type VarietyArchived struct {
	ev.BaseEvent
}

func NewVarietyCreated(id vo.ID) VarietyCreated {
	return VarietyCreated{
		BaseEvent: ev.NewBaseEvent(id, EventVarietyCreated),
	}
}

func NewMaturityUpdated(id vo.ID) MaturityUpdated {
	return MaturityUpdated{
		BaseEvent: ev.NewBaseEvent(id, EventMaturityUpdated),
	}
}
func NewGrowthUpdated(id vo.ID) GrowthUpdated {
	return GrowthUpdated{
		BaseEvent: ev.NewBaseEvent(id, EventGrowthUpdated),
	}
}

func NewSpacingUpdated(id vo.ID) SpacingUpdated {
	return SpacingUpdated{
		BaseEvent: ev.NewBaseEvent(id, EventSpacingUpdated),
	}
}
func NewHarvestUpdated(id vo.ID) HarvestUpdated {
	return HarvestUpdated{
		BaseEvent: ev.NewBaseEvent(id, EventHarvestUpdated),
	}
}

func NewYieldUpdated(id vo.ID) YieldUpdated {
	return YieldUpdated{
		BaseEvent: ev.NewBaseEvent(id, EventYieldUpdated),
	}
}

func NewToleranceUpdated(id vo.ID) ToleranceUpdated {
	return ToleranceUpdated{
		BaseEvent: ev.NewBaseEvent(id, EventToleranceUpdated),
	}
}
func NewVarietyArchived(id vo.ID) VarietyArchived {
	return VarietyArchived{
		BaseEvent: ev.NewBaseEvent(id, EventVarietyArchived),
	}
}
