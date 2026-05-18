package warehouse

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventWarehouseCreated  = "warehouse.created"
	EventWarehouseArchived = "warehouse.archived"
)

type WarehouseCreated struct {
	ev.BaseEvent
}

func NewWarehouseCreated(id vo.ID) WarehouseCreated {
	return WarehouseCreated{ev.NewBaseEvent(id, EventWarehouseCreated)}
}

type WarehouseArchived struct {
	ev.BaseEvent
}

func NewWarehouseArchived(id vo.ID) WarehouseArchived {
	return WarehouseArchived{ev.NewBaseEvent(id, EventWarehouseArchived)}
}
