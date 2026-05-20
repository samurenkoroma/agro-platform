package slot

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Slot struct {
	ev.BaseAggregate
	ID               vo.ID
	ProductionUnitID vo.ID
	Code             string
	Position         *vo.Coordinates
	Status           SlotStatus
	Capacity         Capacity
	Metadata         vo.Metadata
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func New(unitID vo.ID, code string, maxPlants int) (*Slot, error) {

	if maxPlants <= 0 {
		return nil, ErrInvalidCount
	}

	now := time.Now()

	root := &Slot{
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

	root.AddEvent(NewSlotCreated(root.ID))

	return root, nil
}

func (a *Slot) Occupy(count int) error {

	if a.Status == Blocked {
		return ErrSlotBlocked
	}
	if a.Capacity.CurrentPlants+count > a.Capacity.MaxPlants {
		return ErrSlotFull
	}

	a.Capacity.CurrentPlants += count

	if a.Capacity.IsFull() {
		a.Status = Occupied
	}

	a.AddEvent(NewSlotOccupied(a.ID, count))

	return nil
}

func (a *Slot) Release(count int) error {
	if count <= 0 {
		return ErrInvalidCount
	}

	next := a.Capacity.CurrentPlants - count

	if next < 0 {
		return ErrInvalidCount
	}

	a.Capacity.CurrentPlants = next

	if next == 0 {
		a.Status = Available
	}

	a.AddEvent(NewSlotReleased(a.ID, count))

	return nil
}

func (a *Slot) Block() {
	a.Status = Blocked

	a.AddEvent(NewSlotBlocked(a.ID))
}
