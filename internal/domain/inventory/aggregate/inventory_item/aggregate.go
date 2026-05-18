package inventoryitem

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(name string, itemType Type, unit Unit) *Aggregate {
	now := time.Now()

	root := Item{
		ID:        vo.NewID(),
		Name:      name,
		Type:      itemType,
		Unit:      unit,
		Stock:     Stock{},
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewItemCreated(root.ID))

	return a
}

func (a *Aggregate) Receive(amount float64) {
	a.Root.Stock.Available += amount
	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewStockReceived(a.Root.ID, amount))
}

func (a *Aggregate) Reserve(amount float64) error {
	if amount <= 0 {
		return ErrNegativeAmount
	}

	if amount > a.Root.Stock.Available {
		return ErrInsufficientStock
	}

	a.Root.Stock.Available -= amount
	a.Root.Stock.Reserved += amount

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewStockReserved(a.Root.ID, amount))

	return nil
}

func (a *Aggregate) Consume(amount float64) {
	a.Root.Stock.Reserved -= amount

	a.Root.Stock.Consumed += amount

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewStockConsumed(a.Root.ID, amount))
}

func (a *Aggregate) MarkLost(amount float64) {

	a.Root.Stock.Available -= amount

	a.Root.Stock.Lost += amount

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewStockLost(a.Root.ID, amount))
}
