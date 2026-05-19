package substrate

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Substrate struct {
	ev.AggregateRoot
	ID             vo.ID
	Name           string
	Type           SubstrateType
	Reusable       bool
	Status         SubstrateStatus
	VolumeLiters   *float64
	WaterRetention *float64
	Aeration       *float64
	Manufacturer   *string
	BatchID        *vo.ID
	Metadata       vo.Metadata
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func New(name string, t SubstrateType, reusable bool) *Substrate {
	now := time.Now()

	root := &Substrate{
		ID:        vo.NewID(),
		Name:      name,
		Type:      t,
		Reusable:  reusable,
		Status:    Available,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewSubstrateCreated(root.ID))

	return root
}

func (a *Substrate) Use() error {
	if a.Status == Disposed {
		return ErrAlreadyDisposed
	}
	a.Status = InUse
	a.AddEvent(NewSubstrateUsed(a.ID))

	return nil
}

func (a *Substrate) Exhaust() {
	a.Status = Exhausted

	a.AddEvent(NewSubstrateExhausted(a.ID))
}

func (a *Substrate) Recycle() error {
	if !a.Reusable {
		return ErrNotReusable
	}

	a.Status = Recycled

	a.AddEvent(NewSubstrateRecycled(a.ID))

	return nil
}

func (a *Substrate) Dispose() {
	a.Status = Disposed

	a.AddEvent(NewSubstrateDisposed(a.ID))
}
