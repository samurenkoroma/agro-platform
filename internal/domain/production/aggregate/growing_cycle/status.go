package growingcycle

type CycleStatus string

const (
	StatusPlanned    CycleStatus = "planned"
	StatusActive     CycleStatus = "active"
	StatusPaused     CycleStatus = "paused"
	StatusHarvesting CycleStatus = "harvesting"
	StatusCompleted  CycleStatus = "completed"
	StatusFailed     CycleStatus = "failed"
	StatusArchived   CycleStatus = "archived"
)
