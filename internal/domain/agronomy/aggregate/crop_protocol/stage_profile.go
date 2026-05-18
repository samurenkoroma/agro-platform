package cropprotocol

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type StageProfile struct {
	ID          vo.ID
	CropStageID vo.ID
	Climate     ClimateProfile
	Lighting    LightingProfile
	Irrigation  IrrigationProfile
	Nutrition   NutritionProfile
	Water       WaterDemandProfile
	VPD         VPDProfile
}
