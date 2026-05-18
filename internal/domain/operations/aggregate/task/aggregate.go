package task

import (
	"time"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func New(farmID vo.ID, title string) *Aggregate {
	now := time.Now()

	root := Task{
		ID:        vo.NewID(),
		FarmID:    farmID,
		Title:     title,
		Status:    Todo,
		Priority:  Medium,
		Metadata:  vo.NewMetadata(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	a := &Aggregate{
		Root: root,
	}

	a.AddEvent(NewTaskCreated(root.ID))

	return a
}

func (a *Aggregate) Assign(userID vo.ID) {
	a.Root.Assignment = &Assignment{UserID: userID}

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewTaskAssigned(a.Root.ID, userID))
}

func (a *Aggregate) Start() {

	a.Root.Status = InProgress

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewTaskStarted(a.Root.ID))
}

func (a *Aggregate) Complete() {
	now := time.Now()

	a.Root.Status = Done
	a.Root.CompletedAt = &now
	a.Root.UpdatedAt = now

	a.AddEvent(NewTaskCompleted(a.Root.ID))
}

func (a *Aggregate) Cancel() {

	a.Root.Status = Cancelled

	a.Root.UpdatedAt = time.Now()

	a.AddEvent(NewTaskCancelled(a.Root.ID))
}
