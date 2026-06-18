package task_test

import (
	"testing"
	"time"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/task"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

var (
	farmID = vo.NewID()
	userID = vo.NewID()
)

func newTask(t *testing.T) *task.Task {
	t.Helper()
	tsk := task.New(farmID, "Полить рассаду")
	tsk.PullEvents() // сбрасываем TaskCreated
	return tsk
}

// --- New ---

func TestNew_DefaultsAreCorrect(t *testing.T) {
	tsk := task.New(farmID, "Полить рассаду")

	if tsk.ID.IsZero() {
		t.Error("expected non-zero ID")
	}
	if tsk.FarmID != farmID {
		t.Errorf("farmID: got %s, want %s", tsk.FarmID, farmID)
	}
	if tsk.Title != "Полить рассаду" {
		t.Errorf("title: got %s, want Полить рассаду", tsk.Title)
	}
	if tsk.Status != task.Todo {
		t.Errorf("status: got %s, want TODO", tsk.Status)
	}
	if tsk.Priority != task.Medium {
		t.Errorf("priority: got %s, want MEDIUM", tsk.Priority)
	}
	if tsk.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
}

func TestNew_EmitsTaskCreatedEvent(t *testing.T) {
	tsk := task.New(farmID, "Полить рассаду")
	events := tsk.PullEvents()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].EventType() != task.EventTaskCreated {
		t.Errorf("event type: got %s, want %s", events[0].EventType(), task.EventTaskCreated)
	}
	if events[0].AggregateID() != tsk.ID {
		t.Error("event aggregate ID does not match task ID")
	}
}

func TestNew_PullEventsIsDrained(t *testing.T) {
	tsk := task.New(farmID, "test")
	tsk.PullEvents()
	events := tsk.PullEvents()
	if len(events) != 0 {
		t.Errorf("expected 0 events after drain, got %d", len(events))
	}
}

// --- Assign ---

func TestAssign_SetsAssignment(t *testing.T) {
	tsk := newTask(t)
	tsk.Assign(userID)

	if tsk.Assignment == nil {
		t.Fatal("expected assignment to be set")
	}
	if tsk.Assignment.UserID != userID {
		t.Errorf("userID: got %s, want %s", tsk.Assignment.UserID, userID)
	}
}

func TestAssign_EmitsTaskAssignedEvent(t *testing.T) {
	tsk := newTask(t)
	tsk.Assign(userID)

	events := tsk.PullEvents()
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].EventType() != task.EventTaskAssigned {
		t.Errorf("event type: got %s, want %s", events[0].EventType(), task.EventTaskAssigned)
	}
}

func TestAssign_UpdatesUpdatedAt(t *testing.T) {
	tsk := newTask(t)
	before := tsk.UpdatedAt
	time.Sleep(time.Millisecond)
	tsk.Assign(userID)

	if !tsk.UpdatedAt.After(before) {
		t.Error("UpdatedAt should be updated after Assign")
	}
}

func TestAssign_OverwritesPreviousAssignment(t *testing.T) {
	tsk := newTask(t)
	otherUser := vo.NewID()
	tsk.Assign(userID)
	tsk.Assign(otherUser)

	if tsk.Assignment.UserID != otherUser {
		t.Errorf("expected assignment to be overwritten to %s", otherUser)
	}
}

// --- Start ---

func TestStart_ChangesStatusToInProgress(t *testing.T) {
	tsk := newTask(t)
	tsk.Start()

	if tsk.Status != task.InProgress {
		t.Errorf("status: got %s, want IN_PROGRESS", tsk.Status)
	}
}

func TestStart_EmitsTaskStartedEvent(t *testing.T) {
	tsk := newTask(t)
	tsk.Start()

	events := tsk.PullEvents()
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].EventType() != task.EventTaskStarted {
		t.Errorf("event type: got %s, want %s", events[0].EventType(), task.EventTaskStarted)
	}
}

// --- Complete ---

func TestComplete_ChangesStatusToDone(t *testing.T) {
	tsk := newTask(t)
	tsk.Start()
	tsk.PullEvents()
	tsk.Complete()

	if tsk.Status != task.Done {
		t.Errorf("status: got %s, want DONE", tsk.Status)
	}
}

func TestComplete_SetsCompletedAt(t *testing.T) {
	tsk := newTask(t)
	tsk.Complete()

	if tsk.CompletedAt == nil {
		t.Fatal("CompletedAt should be set after Complete")
	}
	if tsk.CompletedAt.IsZero() {
		t.Error("CompletedAt should not be zero")
	}
}

func TestComplete_EmitsTaskCompletedEvent(t *testing.T) {
	tsk := newTask(t)
	tsk.Complete()

	events := tsk.PullEvents()
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].EventType() != task.EventTaskCompleted {
		t.Errorf("event type: got %s, want %s", events[0].EventType(), task.EventTaskCompleted)
	}
}

// --- Cancel ---

func TestCancel_ChangesStatusToCancelled(t *testing.T) {
	tsk := newTask(t)
	tsk.Cancel()

	if tsk.Status != task.Cancelled {
		t.Errorf("status: got %s, want CANCELLED", tsk.Status)
	}
}

func TestCancel_EmitsTaskCancelledEvent(t *testing.T) {
	tsk := newTask(t)
	tsk.Cancel()

	events := tsk.PullEvents()
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].EventType() != task.EventTaskCancelled {
		t.Errorf("event type: got %s, want %s", events[0].EventType(), task.EventTaskCancelled)
	}
}

// --- State transitions ---

func TestStatusTransition_TodoToInProgressToDone(t *testing.T) {
	tsk := newTask(t)

	tsk.Start()
	if tsk.Status != task.InProgress {
		t.Errorf("after Start: got %s, want IN_PROGRESS", tsk.Status)
	}

	tsk.Complete()
	if tsk.Status != task.Done {
		t.Errorf("after Complete: got %s, want DONE", tsk.Status)
	}
}

func TestStatusTransition_TodoToCancelled(t *testing.T) {
	tsk := newTask(t)
	tsk.Cancel()

	if tsk.Status != task.Cancelled {
		t.Errorf("got %s, want CANCELLED", tsk.Status)
	}
}

// --- Priority ---

func TestPriority_CanBeSetToCritical(t *testing.T) {
	tsk := newTask(t)
	tsk.Priority = task.Critical

	if tsk.Priority != task.Critical {
		t.Errorf("priority: got %s, want CRITICAL", tsk.Priority)
	}
}

// --- Optional fields ---

func TestTask_DescriptionCanBeSet(t *testing.T) {
	tsk := newTask(t)
	desc := "Проверить датчики"
	tsk.Description = &desc

	if tsk.Description == nil || *tsk.Description != desc {
		t.Error("description not set correctly")
	}
}

func TestTask_DueDateCanBeSet(t *testing.T) {
	tsk := newTask(t)
	due := time.Now().Add(24 * time.Hour)
	tsk.DueDate = &due

	if tsk.DueDate == nil {
		t.Fatal("DueDate should be set")
	}
	if !tsk.DueDate.Equal(due) {
		t.Error("DueDate mismatch")
	}
}
