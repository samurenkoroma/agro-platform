package cropprotocol

import (
	"time"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(cropID vo.ID, name string, method gc.GrowingMethod) *Aggregate {
	now := time.Now()

	root := CropProtocol{
		ID:            vo.NewID(),
		CropID:        cropID,
		Name:          name,
		GrowingMethod: method,
		StageProfiles: make(
			[]StageProfile,
			0,
		),
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewProtocolCreated(root.ID))

	return a
}

func (a *Aggregate) AddStage(stage StageProfile) {
	a.Root.StageProfiles = append(a.Root.StageProfiles, stage)

	a.AddEvent(NewStageAdded(a.Root.ID, stage.ID))
}
