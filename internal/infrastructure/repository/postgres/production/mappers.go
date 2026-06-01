package production

import growingcycle "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/growing_cycle"

func scanGrowingCycle(root *growingcycle.GrowingCycle) []any {
	return []any{
		&root.ID,
		&root.FarmID,
		&root.CropID,
		&root.VarietyID,
		&root.ProtocolID,
		&root.Name,
		&root.Code,
		&root.Method,
		&root.Status,
		&root.Stage,
		&root.StartedAt,
		&root.CompletedAt,
		&root.ExpectedHarvestAt,
		&root.CreatedAt,
		&root.UpdatedAt,
	}
}
