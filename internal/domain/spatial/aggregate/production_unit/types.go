package productionunit

type ProductionUnitType string

const (
	// ========== ПОЛЕВЫЕ ТИПЫ (Open Field) ==========
	Field ProductionUnitType = "FIELD"
	Plot  ProductionUnitType = "PLOT"
	Block ProductionUnitType = "BLOCK"
	Bed   ProductionUnitType = "BED"
	Row   ProductionUnitType = "ROW"

	// ========== ТЕПЛИЧНЫЕ ТИПЫ (Greenhouse) ==========
	Greenhouse     ProductionUnitType = "GREENHOUSE"
	GreenhouseZone ProductionUnitType = "GREENHOUSE_ZONE"

	// ========== КОНТЕЙНЕРНЫЕ ТИПЫ (Container-based) ==========
	Container     ProductionUnitType = "CONTAINER"
	Rack          ProductionUnitType = "RACK"
	Shelf         ProductionUnitType = "SHELF"
	VerticalTower ProductionUnitType = "VERTICAL_TOWER"

	// ========== ПОСАДОЧНЫЕ МЕСТА (Planting Slots) ==========
	Slot ProductionUnitType = "SLOT"
	Pot  ProductionUnitType = "POT"
	Tray ProductionUnitType = "TRAY"

	// ========== ГИДРОПОННЫЕ СИСТЕМЫ (Hydroponic) ==========
	NFTChannel  ProductionUnitType = "NFT_CHANNEL"
	DWCTank     ProductionUnitType = "DWC_TANK"
	AeroChamber ProductionUnitType = "AEROPONIC_CHAMBER"

	// ========== СЛУЖЕБНЫЕ ТИПЫ (Utility) ==========
	Reservoir ProductionUnitType = "RESERVOIR"
	Storage   ProductionUnitType = "STORAGE"
)
