package cropprotocol

import (
	"time"

	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type CropProtocol struct {
	ev.BaseAggregate
	ID            vo.ID
	CropID        vo.ID
	Name          string
	GrowingMethod gc.GrowingMethod
	Description   *string
	StageProfiles []StageProfile
	Metadata      vo.Metadata
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func New(cropID vo.ID, name string, method gc.GrowingMethod) *CropProtocol {
	now := time.Now()

	root := &CropProtocol{
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

	root.AddEvent(NewProtocolCreated(root.ID))

	return root
}

func (a *CropProtocol) AddStage(stage StageProfile) {
	a.StageProfiles = append(a.StageProfiles, stage)

	a.AddEvent(NewStageAdded(a.ID, stage.ID))
}
