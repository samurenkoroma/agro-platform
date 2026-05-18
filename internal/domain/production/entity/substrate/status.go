package substrate

type SubstrateStatus string

const (
	Available SubstrateStatus = "AVAILABLE"
	InUse     SubstrateStatus = "IN_USE"
	Exhausted SubstrateStatus = "EXHAUSTED"
	Disposed  SubstrateStatus = "DISPOSED"
	Recycled  SubstrateStatus = "RECYCLED"
)
