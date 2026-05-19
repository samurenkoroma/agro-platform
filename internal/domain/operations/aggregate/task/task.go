package task

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	oe "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
)

type Task struct {
	ID               vo.ID
	Title            string
	Description      *string
	OperationType    *oe.OperationType
	FarmID           vo.ID
	ProductionUnitID *vo.ID
	GrowingCycleID   *vo.ID
	PlantID          *vo.ID
	Assignment       *Assignment
	Status           Status
	Priority         Priority
	DueDate          *time.Time
	CompletedAt      *time.Time
	Metadata         vo.Metadata
	CreatedAt        time.Time
	UpdatedAt        time.Time
	ArchivedAt       *time.Time
}

type Aggregate struct {
	ev.AggregateRoot
	Root Task
}
