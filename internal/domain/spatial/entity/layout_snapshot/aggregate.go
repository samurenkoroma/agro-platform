package layoutsnapshot

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pus "github.com/samurenkoroma/agro-platform/internal/domain/spatial/entity/production_unit_snapshot"
)

func New(
	farmID vo.ID,

	version int,

	createdBy vo.ID,

	description *string,
) (
	*Aggregate,
	error,
) {

	if version <= 0 {
		return nil,
			ErrInvalidVersion
	}

	root := LayoutSnapshot{
		ID: vo.NewID(),

		FarmID: farmID,

		Version: version,

		Description: description,

		CreatedBy: createdBy,

		CreatedAt: time.Now(),

		Units: make(
			[]pus.ProductionUnitSnapshot,
			0,
		),
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(
		NewLayoutCreated(
			root.ID,
		),
	)

	return a, nil
}

func (
	a *Aggregate,
) AddUnit(
	unit pus.ProductionUnitSnapshot,
) error {

	for _, existing := range a.Root.Units {

		if existing.
			OriginalUnitID ==
			unit.
				OriginalUnitID {

			return ErrDuplicateUnit
		}
	}

	a.Root.Units =
		append(
			a.Root.Units,
			unit,
		)

	a.AddEvent(
		NewUnitCaptured(
			a.Root.ID,
			unit.OriginalUnitID,
		),
	)

	return nil
}
