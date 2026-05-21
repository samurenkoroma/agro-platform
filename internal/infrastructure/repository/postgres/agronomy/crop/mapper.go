package crop

import (
	entity "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
)

func scanTarget(root *entity.Crop) []any {
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
