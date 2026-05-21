package agronomy

import (
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	protocol "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop_protocol"
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

		&root.Harvest,
		&root.Yield,

		&root.Tolerance,

		&root.Metadata,

		&root.CreatedAt,
		&root.UpdatedAt,
		&root.ArchivedAt,
	}
}

func scanCropProtocol(root *protocol.CropProtocol) []any {
	return []any{
		&root.ID,
		&root.CropID,
		&root.Name,
		&root.GrowingMethod,
		&root.Description,
		&root.Metadata,
		&root.CreatedAt,
		&root.UpdatedAt,
		&root.ArchivedAt,
	}
}

func scanStage(root *protocol.StageProfile) []any {
	return []any{
		&root.ID,
		&root.CropStageID,

		&root.Climate,
		&root.Lighting,
		&root.Irrigation,
		&root.Nutrition,
		&root.Water,
		&root.VPD,
	}
}
