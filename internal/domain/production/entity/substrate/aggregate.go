package substrate

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(name string, t SubstrateType, reusable bool) *Aggregate {
	now := time.Now()

	root := Substrate{
		ID:        vo.NewID(),
		Name:      name,
		Type:      t,
		Reusable:  reusable,
		Status:    Available,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewSubstrateCreated(root.ID))

	return a
}

func (a *Aggregate) Use() error {
	if a.Root.Status == Disposed {
		return ErrAlreadyDisposed
	}

	a.Root.Status = InUse

	a.AddEvent(NewSubstrateUsed(a.Root.ID))

	return nil
}

func (a *Aggregate) Exhaust() {
	a.Root.Status = Exhausted

	a.AddEvent(NewSubstrateExhausted(a.Root.ID))
}

func (a *Aggregate) Recycle() error {
	if !a.Root.Reusable {
		return ErrNotReusable
	}

	a.Root.Status = Recycled

	a.AddEvent(NewSubstrateRecycled(a.Root.ID))

	return nil
}

func (a *Aggregate) Dispose() {
	a.Root.Status = Disposed

	a.AddEvent(NewSubstrateDisposed(a.Root.ID))
}
