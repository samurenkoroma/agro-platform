package operationevent

type OperationType string

const (
	Plowed            OperationType = "PLOWED"
	Sowed             OperationType = "SOWED"
	Transplanted      OperationType = "TRANSPLANTED"
	Irrigated         OperationType = "IRRIGATED"
	Fertilized        OperationType = "FERTILIZED"
	Sprayed           OperationType = "SPRAYED"
	Harvested         OperationType = "HARVESTED"
	PlantDiscarded    OperationType = "PLANT_DISCARDED"
	PHAdjusted        OperationType = "PH_ADJUSTED"
	ECAdjusted        OperationType = "EC_ADJUSTED"
	ReservoirRefilled OperationType = "RESERVOIR_REFILLED"
	LightChanged      OperationType = "LIGHT_CHANGED"
	SensorAlert       OperationType = "SENSOR_ALERT"
	LayoutChanged     OperationType = "LAYOUT_CHANGED"
)
