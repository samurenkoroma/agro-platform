package slot

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(unitID vo.ID, code string, maxPlants int) (*Aggregate, error) {

	if maxPlants <= 0 {
		return nil, ErrInvalidCount
	}

	now := time.Now()

	root := Slot{
		ID:               vo.NewID(),
		ProductionUnitID: unitID,
		Code:             code,
		Status:           Available,
		Capacity: Capacity{
			MaxPlants: maxPlants,
		},
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewSlotCreated(root.ID))

	return a, nil
}

func (a *Aggregate) Occupy(count int) error {

	if a.Root.Status == Blocked {
		return ErrSlotBlocked
	}
	if a.Root.Capacity.CurrentPlants+count > a.Root.Capacity.MaxPlants {
		return ErrSlotFull
	}

	a.Root.Capacity.CurrentPlants += count

	if a.Root.Capacity.IsFull() {
		a.Root.Status = Occupied
	}

	a.AddEvent(NewSlotOccupied(a.Root.ID, count))

	return nil
}

func (a *Aggregate) Release(count int) error {
	if count <= 0 {
		return ErrInvalidCount
	}

	next := a.Root.Capacity.CurrentPlants - count

	if next < 0 {
		return ErrInvalidCount
	}

	a.Root.Capacity.CurrentPlants = next

	if next == 0 {
		a.Root.Status = Available
	}

	a.AddEvent(NewSlotReleased(a.Root.ID, count))

	return nil
}

func (a *Aggregate) Block() {
	a.Root.Status = Blocked

	a.AddEvent(NewSlotBlocked(a.Root.ID))
}
