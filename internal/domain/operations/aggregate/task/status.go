package task

type Status string

const (
	Todo       Status = "TODO"
	InProgress Status = "IN_PROGRESS"
	Done       Status = "DONE"
	Cancelled  Status = "CANCELLED"
)
