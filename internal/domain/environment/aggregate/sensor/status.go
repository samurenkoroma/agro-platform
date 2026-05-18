package sensor

type Status string

const (
	Online Status = "ONLINE"

	Offline Status = "OFFLINE"

	Fault Status = "FAULT"

	Maintenance Status = "MAINTENANCE"
)
