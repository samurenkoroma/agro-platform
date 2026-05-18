package task

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventTaskCreated   = "task.created"
	EventTaskAssigned  = "task.assigned"
	EventTaskStarted   = "task.started"
	EventTaskCompleted = "task.completed"
	EventTaskCancelled = "task.cancelled"
)

type TaskCreated struct {
	ev.BaseEvent
}

func NewTaskCreated(id vo.ID) TaskCreated {
	return TaskCreated{
		BaseEvent: ev.NewBaseEvent(id, EventTaskCreated),
	}
}

type TaskAssigned struct {
	ev.BaseEvent
	userID vo.ID
}

func NewTaskAssigned(id vo.ID, userID vo.ID) TaskAssigned {
	return TaskAssigned{
		BaseEvent: ev.NewBaseEvent(id, EventTaskAssigned),
		userID:    userID,
	}
}

type TaskStarted struct {
	ev.BaseEvent
}

func NewTaskStarted(id vo.ID) TaskStarted {
	return TaskStarted{
		BaseEvent: ev.NewBaseEvent(id, EventTaskStarted),
	}
}

type TaskCompleted struct {
	ev.BaseEvent
}

func NewTaskCompleted(id vo.ID) TaskCompleted {
	return TaskCompleted{
		BaseEvent: ev.NewBaseEvent(id, EventTaskCompleted),
	}
}

type TaskCancelled struct {
	ev.BaseEvent
}

func NewTaskCancelled(id vo.ID) TaskCancelled {
	return TaskCancelled{
		BaseEvent: ev.NewBaseEvent(id, EventTaskCancelled),
	}
}
