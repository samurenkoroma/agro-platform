package productionunit

type ProductionUnitType string

const (
	// open field
	Field ProductionUnitType = "FIELD"
	Block ProductionUnitType = "BLOCK"
	Bed   ProductionUnitType = "BED"
	Row   ProductionUnitType = "ROW"

	// greenhouse
	Greenhouse     ProductionUnitType = "GREENHOUSE"
	GreenhouseZone ProductionUnitType = "GREENHOUSE_ZONE"

	// containers
	Container ProductionUnitType = "CONTAINER"
	Pot       ProductionUnitType = "POT"
	Tray      ProductionUnitType = "TRAY"

	// hydro
	NFTChannel  ProductionUnitType = "NFT_CHANNEL"
	DWCTank     ProductionUnitType = "DWC_TANK"
	AeroChamber ProductionUnitType = "AEROPONIC_CHAMBER"

	// vertical
	Rack          ProductionUnitType = "RACK"
	Shelf         ProductionUnitType = "SHELF"
	VerticalTower ProductionUnitType = "VERTICAL_TOWER"

	Slot ProductionUnitType = "SLOT"

	Reservoir ProductionUnitType = "RESERVOIR"

	Storage ProductionUnitType = "STORAGE"
	Room    ProductionUnitType = "room"
	Zone    ProductionUnitType = "zone"
	Tank    ProductionUnitType = "tank"
)
