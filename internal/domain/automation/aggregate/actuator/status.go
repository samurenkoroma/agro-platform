package actuator

type Status string

const (
	Enabled     Status = "ENABLED"
	Disabled    Status = "DISABLED"
	Fault       Status = "FAULT"
	Maintenance Status = "MAINTENANCE"
)
