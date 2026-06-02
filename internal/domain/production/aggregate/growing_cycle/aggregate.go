package growingcycle

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type GrowingCycle struct {
	ev.BaseAggregate
	ID vo.ID

	FarmID vo.ID

	CropID    vo.ID
	VarietyID *vo.ID

	ProtocolID *vo.ID

	Name string
	Code string

	Method string

	Status CycleStatus
	Stage  CycleStage

	StartedAt   *time.Time
	CompletedAt *time.Time

	ExpectedHarvestAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(orgId, cropId vo.ID, varietyId, protocolId *vo.ID, name, code, method string, expected *time.Time) *GrowingCycle {
	now := time.Now()
	return &GrowingCycle{
		ID:                vo.NewID(),
		FarmID:            orgId,
		CropID:            cropId,
		VarietyID:         varietyId,
		ProtocolID:        protocolId,
		Name:              name,
		Code:              code,
		Method:            method,
		Status:            StatusPlanned,
		Stage:             StagePlanning,
		ExpectedHarvestAt: expected,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}
