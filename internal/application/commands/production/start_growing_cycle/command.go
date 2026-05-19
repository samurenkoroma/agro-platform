package startgrowingcycle

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const CommandName = "production.start_growing_cycle"

type Command struct {
	FarmID            vo.ID
	CropID            vo.ID
	VarietyID         *vo.ID
	ProductionUnitID  vo.ID
	ExpectedHarvestAt *time.Time
}

func (Command) CommandName() string {
	return CommandName
}
