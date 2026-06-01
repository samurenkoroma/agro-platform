package cropprotocol

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CropProtocol struct {
	ev.BaseAggregate
	ID            vo.ID
	CropID        vo.ID
	Name          string
	Description   *string
	StageProfiles []StageProfile
	Metadata      vo.Metadata
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ArchivedAt    *time.Time
}

func (a *CropProtocol) AddStage(stage StageProfile) {
	a.StageProfiles = append(a.StageProfiles, stage)

	a.AddEvent(NewStageAdded(a.ID, stage.ID))
}
