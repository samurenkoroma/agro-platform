package inventoryitem

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventItemCreated   = "inventory.item.created"
	EventStockReceived = "inventory.stock.received"
	EventStockReserved = "inventory.stock.reserved"
	EventStockConsumed = "inventory.stock.consumed"
	EventStockLost     = "inventory.stock.lost"
	EventItemArchived  = "inventory.item.archived"
)

type ItemCreated struct {
	ev.BaseEvent
}

func NewItemCreated(id vo.ID) ItemCreated {
	return ItemCreated{BaseEvent: ev.NewBaseEvent(id, EventItemCreated)}
}

type StockReceived struct {
	ev.BaseEvent
	amount float64
}

func NewStockReceived(id vo.ID, amount float64) StockReceived {
	return StockReceived{BaseEvent: ev.NewBaseEvent(id, EventStockReceived), amount: amount}
}

type StockReserved struct {
	ev.BaseEvent
	amount float64
}

func NewStockReserved(id vo.ID, amount float64) StockReserved {
	return StockReserved{BaseEvent: ev.NewBaseEvent(id, EventStockReserved), amount: amount}
}

type StockConsumed struct {
	ev.BaseEvent
	amount float64
}

func NewStockConsumed(id vo.ID, amount float64) StockConsumed {
	return StockConsumed{BaseEvent: ev.NewBaseEvent(id, EventStockConsumed), amount: amount}
}

type StockLost struct {
	ev.BaseEvent
	amount float64
}

func NewStockLost(id vo.ID, amount float64) StockLost {
	return StockLost{BaseEvent: ev.NewBaseEvent(id, EventStockLost), amount: amount}
}

type ItemArchived struct {
	ev.BaseEvent
}

func NewItemArchived(id vo.ID) ItemArchived {
	return ItemArchived{BaseEvent: ev.NewBaseEvent(id, EventItemArchived)}
}
