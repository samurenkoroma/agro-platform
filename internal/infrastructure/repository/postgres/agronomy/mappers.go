package agronomy

import (
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/variety"
)

func scanCrop(root *crop.Crop) []any {
	return []any{
		&root.ID,
		&root.Name,
		&root.ScientificName,
		&root.Category,
		&root.Metadata,
		&root.CreatedAt,
		&root.UpdatedAt,
	}
}

func scanVariety(root *variety.Variety) []any {
	return []any{
		&root.ID,
		&root.CropID,
		&root.Name,
		&root.Breeder,
		&root.Maturity,
		&root.Growth,
		&root.Spacing,
		&root.Tolerance,
		&root.Metadata,
		&root.CreatedAt,
		&root.UpdatedAt,
	}
}
