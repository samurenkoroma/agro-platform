package plant

type PlantStatus string

const (
	Germinating PlantStatus = "GERMINATING"

	Active PlantStatus = "ACTIVE"

	Transplanted PlantStatus = "TRANSPLANTED"

	Stressed PlantStatus = "STRESSED"

	Diseased PlantStatus = "DISEASED"

	Harvesting PlantStatus = "HARVESTING"

	Harvested PlantStatus = "HARVESTED"

	Discarded PlantStatus = "DISCARDED"

	Dead PlantStatus = "DEAD"
)
