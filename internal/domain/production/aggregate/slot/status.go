package slot

type SlotStatus string

const (
	Available SlotStatus = "AVAILABLE"

	Occupied SlotStatus = "OCCUPIED"

	Reserved SlotStatus = "RESERVED"

	Blocked SlotStatus = "BLOCKED"

	Disabled SlotStatus = "DISABLED"
)
