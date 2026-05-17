package plant

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type PlantStage struct {
	StageID vo.ID

	ChangedAt time.Time
}

func (a *Aggregate) MoveStage(
	stageID vo.ID,
) error {

	if a.Root.Status ==
		Harvested {

		return ErrAlreadyHarvested
	}

	if a.Root.Status ==
		Discarded {

		return ErrAlreadyDiscarded
	}

	if a.Root.Status ==
		Dead {

		return ErrPlantDead
	}

	now := time.Now()

	a.Root.CurrentStageID =
		&stageID

	a.Root.UpdatedAt =
		now

	a.AddEvent(
		NewPlantStageChanged(
			a.Root.ID,
			stageID,
		),
	)

	return nil
}
