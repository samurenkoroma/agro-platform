package production

import (
	gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"
	harvest "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/harvest_batch"
	"github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/plant"
	yieldbatch "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/yield_batch"
)

func scanGrowingCycle(root *gc.GrowingCycle) []any {
	return []any{
		&root.ID,
		&root.FarmID,
		&root.CropID,
		&root.ProductionUnitID,
		&root.Method,
		&root.Status,
		&root.Metadata,
		&root.CreatedAt,
		&root.UpdatedAt,
		&root.ArchivedAt,
	}
}

func scanPlant(root *plant.Plant) []any {
	return []any{
		&root.ID,
		&root.GrowingCycleID,
		&root.CropID,
		&root.VarietyID,
		&root.ProductionUnitID,
		&root.SlotID,
		&root.SubstrateID,
		&root.Status,
		&root.Health,
		&root.CurrentStageID,
		&root.PlantedAt,
		&root.TransplantedAt,
		&root.HarvestedAt,
		&root.DiscardedAt,
		&root.Metadata,
		&root.CreatedAt,
		&root.UpdatedAt,
	}
}

func scanHarvestBatch(root *harvest.HarvestBatch) []any {
	return []any{
		&root.ID,
		&root.GrowingCycleID,
		&root.ProductionUnitID,
		&root.Quantity,
		&root.HarvestedArea,
		&root.Grade,
		&root.Marketable,
		&root.Notes,
		&root.HarvestedAt,
		&root.Metadata,
		&root.CreatedAt,
	}
}

func scanYieldBatch(root *yieldbatch.YieldBatch) []any {
	return []any{
		&root.ID,
		&root.GrowingCycleID,
		&root.PlantID,
		&root.Quantity,
		&root.FruitCount,
		&root.Grade,
		&root.Marketable,
		&root.Notes,
		&root.HarvestedAt,
		&root.Metadata,
		&root.CreatedAt,
	}
}
