package task

import (
	"time"

	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/aggregate"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	oe "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
)

type Task struct {
	ev.AggregateRoot
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

func New(farmID vo.ID, title string) *Task {
	now := time.Now()

	root := &Task{
		ID:        vo.NewID(),
		FarmID:    farmID,
		Title:     title,
		Status:    Todo,
		Priority:  Medium,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	root.AddEvent(NewTaskCreated(root.ID))

	return root
}

func (a *Task) Assign(userID vo.ID) {
	a.Assignment = &Assignment{UserID: userID}

	a.UpdatedAt = time.Now()

	a.AddEvent(NewTaskAssigned(a.ID, userID))
}

func (a *Task) Start() {

	a.Status = InProgress

	a.UpdatedAt = time.Now()

	a.AddEvent(NewTaskStarted(a.ID))
}

func (a *Task) Complete() {
	now := time.Now()

	a.Status = Done
	a.CompletedAt = &now
	a.UpdatedAt = now

	a.AddEvent(NewTaskCompleted(a.ID))
}

func (a *Task) Cancel() {

	a.Status = Cancelled

	a.UpdatedAt = time.Now()

	a.AddEvent(NewTaskCancelled(a.ID))
}
