package movement

type Type string

const (
	Inbound Type = "INBOUND"

	Outbound Type = "OUTBOUND"

	Transfer Type = "TRANSFER"

	Reserve Type = "RESERVE"

	Consume Type = "CONSUME"

	Loss Type = "LOSS"

	Correction Type = "CORRECTION"
)
