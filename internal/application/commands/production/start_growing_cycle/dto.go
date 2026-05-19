package startgrowingcycle

type DTO struct {
	FarmID            string  `json:"farmId"`
	CropID            string  `json:"cropId"`
	VarietyID         *string `json:"varietyId"`
	ProductionUnitID  string  `json:"productionUnitId"`
	ExpectedHarvestAt *string `json:"expectedHarvestAt"`
}
