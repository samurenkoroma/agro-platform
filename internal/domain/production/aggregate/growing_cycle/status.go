package growingcycle

type GrowingStatus string

const (
	Planned GrowingStatus = "PLANNED"

	Active GrowingStatus = "ACTIVE"

	Harvesting GrowingStatus = "HARVESTING"

	Paused GrowingStatus = "PAUSED"

	Harvested GrowingStatus = "HARVESTED"

	Failed GrowingStatus = "FAILED"

	Archived GrowingStatus = "ARCHIVED"
)
