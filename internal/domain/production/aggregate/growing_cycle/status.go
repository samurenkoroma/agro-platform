package growingcycle

type GrowingStatus string

const (
	Planned GrowingStatus = "PLANNED"

	Active GrowingStatus = "ACTIVE"

	Paused GrowingStatus = "PAUSED"

	Harvested GrowingStatus = "HARVESTED"

	Failed GrowingStatus = "FAILED"

	Archived GrowingStatus = "ARCHIVED"
)
