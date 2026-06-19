package task_test

import (
	"context"
	"testing"

	task2 "github.com/samurenkoroma/agro-platform/internal/application/commands/operations/task"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/task"
	opsrepo "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	inmemops "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/operations"
	"github.com/samurenkoroma/agro-platform/internal/testutil"
)

func newHandler() (*task2.Handler, opsrepo.OperationsProvider) {
	p := inmemops.NewProvider()
	uow := &testutil.FakeUoW{Provider: p}
	return task2.NewTaskHandler(uow), p
}

// --- Create ---

func TestCreateTask_ReturnsID(t *testing.T) {
	h, _ := newHandler()
	result, err := h.Create(testutil.OrgCtx(), &task2.CreateTaskCommand{
		Title:    "Полить рассаду",
		Priority: "HIGH",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected result ID, got nil")
	}
}

func TestCreateTask_FailsWithoutOrgID(t *testing.T) {
	h, _ := newHandler()
	_, err := h.Create(context.Background(), &task2.CreateTaskCommand{Title: "test"})
	if err == nil {
		t.Fatal("expected error when organization_id is missing")
	}
}

func TestCreateTask_FailsWithWrongPayloadType(t *testing.T) {
	h, _ := newHandler()
	_, err := h.Create(testutil.OrgCtx(), "not a command")
	if err == nil {
		t.Fatal("expected error for invalid payload type")
	}
}

func TestCreateTask_PersistsToRepository(t *testing.T) {
	h, p := newHandler()
	farmID := vo.NewID()
	ctx := testutil.OrgCtxWithID(farmID)

	_, err := h.Create(ctx, &task2.CreateTaskCommand{Title: "Удобрить", Priority: "MEDIUM"})
	if err != nil {
		t.Fatalf("create: %v", err)
	}

	tasks, err := p.Tasks().List(ctx, opsrepo.TaskFilter{FarmID: farmID})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Title != "Удобрить" {
		t.Errorf("title: got %s, want Удобрить", tasks[0].Title)
	}
}

func TestCreateTask_DefaultStatusIsTodo(t *testing.T) {
	h, p := newHandler()
	farmID := vo.NewID()
	ctx := testutil.OrgCtxWithID(farmID)

	h.Create(ctx, &task2.CreateTaskCommand{Title: "test"})

	tasks, _ := p.Tasks().List(ctx, opsrepo.TaskFilter{FarmID: farmID})
	if tasks[0].Status != task.Todo {
		t.Errorf("status: got %s, want TODO", tasks[0].Status)
	}
}

func TestCreateTask_WithOptionalFields(t *testing.T) {
	h, p := newHandler()
	farmID := vo.NewID()
	ctx := testutil.OrgCtxWithID(farmID)

	cycleID := vo.NewID().String()
	unitID := vo.NewID().String()
	desc := "Описание"

	h.Create(ctx, &task2.CreateTaskCommand{
		Title:            "С полями",
		Description:      &desc,
		GrowingCycleID:   &cycleID,
		ProductionUnitID: &unitID,
		Priority:         "LOW",
	})

	tasks, _ := p.Tasks().List(ctx, opsrepo.TaskFilter{FarmID: farmID})
	tsk := tasks[0]

	if tsk.Description == nil || *tsk.Description != desc {
		t.Error("description mismatch")
	}
	if tsk.GrowingCycleID == nil {
		t.Error("GrowingCycleID should be set")
	}
	if tsk.ProductionUnitID == nil {
		t.Error("ProductionUnitID should be set")
	}
}
