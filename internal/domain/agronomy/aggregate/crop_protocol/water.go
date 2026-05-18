package cropprotocol

type WaterDemandProfile struct {
	LitersPerPlantDay       *float64
	LitersPerSquareMeterDay *float64
	SupportsRecirculation   bool
}
