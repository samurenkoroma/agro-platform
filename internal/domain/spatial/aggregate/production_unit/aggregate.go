package productionunit

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(
	farmID vo.ID,

	unitType ProductionUnitType,

	name string,
) (
	*Aggregate,
	error,
) {

	if name == "" {
		return nil,
			ErrInvalidName
	}

	root := ProductionUnit{
		ID: vo.NewID(),

		FarmID: farmID,

		Type: unitType,

		Name: name,

		Metadata: vo.NewMetadata(),

		CreatedAt: time.Now(),

		UpdatedAt: time.Now(),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(
		NewProductionUnitCreated(
			root.ID,
		),
	)

	return a, nil
}

func (a *Aggregate) AttachTo(
	parentID vo.ID,
) error {

	if a.Root.ParentID != nil {

		return ErrAlreadyHasParent
	}

	a.Root.ParentID =
		&parentID

	a.AddEvent(
		NewProductionUnitAttached(
			a.Root.ID,
			parentID,
		),
	)

	return nil
}

func (a *Aggregate) AddCapability(
	c Capability,
) {

	for _, v := range a.Root.Capabilities {

		if v == c {
			return
		}
	}

	a.Root.Capabilities =
		append(
			a.Root.Capabilities,
			c,
		)
}
