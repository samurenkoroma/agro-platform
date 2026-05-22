package growingcycle

import gc "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"

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
