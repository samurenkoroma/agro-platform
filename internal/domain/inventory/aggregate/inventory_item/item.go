package inventoryitem

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Item struct {
	ev.BaseAggregate
	ID          vo.ID
	Name        string
	Type        Type
	Unit        Unit
	WarehouseID *vo.ID
	Stock       Stock
	Metadata    vo.Metadata
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ArchivedAt  *time.Time
}

func New(name string, itemType Type, unit Unit) *Item {
	now := time.Now()

	root := &Item{
		ID:        vo.NewID(),
		Name:      name,
		Type:      itemType,
		Unit:      unit,
		Stock:     Stock{},
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewItemCreated(root.ID))

	return root
}

func (a *Item) Receive(amount float64) {
	a.Stock.Available += amount
	a.UpdatedAt = time.Now()

	a.AddEvent(NewStockReceived(a.ID, amount))
}

func (a *Item) Reserve(amount float64) error {
	if amount <= 0 {
		return ErrNegativeAmount
	}

	if amount > a.Stock.Available {
		return ErrInsufficientStock
	}

	a.Stock.Available -= amount
	a.Stock.Reserved += amount

	a.UpdatedAt = time.Now()

	a.AddEvent(NewStockReserved(a.ID, amount))

	return nil
}

func (a *Item) Consume(amount float64) {
	a.Stock.Reserved -= amount
	a.Stock.Consumed += amount
	a.UpdatedAt = time.Now()

	a.AddEvent(NewStockConsumed(a.ID, amount))
}

func (a *Item) MarkLost(amount float64) {
	a.Stock.Available -= amount
	a.Stock.Lost += amount
	a.UpdatedAt = time.Now()

	a.AddEvent(NewStockLost(a.ID, amount))
}
