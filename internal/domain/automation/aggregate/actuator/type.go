package actuator

type Type string

const (
	Pump   Type = "PUMP"
	Valve  Type = "VALVE"
	Light  Type = "LIGHT"
	Fan    Type = "FAN"
	Heater Type = "HEATER"
	Cooler Type = "COOLER"
	Doser  Type = "DOSER"
	Mister Type = "MISTER"
)
