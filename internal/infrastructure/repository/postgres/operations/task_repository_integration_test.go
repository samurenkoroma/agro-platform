//go:build integration

package operations_test

import (
	"context"
	"testing"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/task"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pgops "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/operations"
	dbtest "github.com/samurenkoroma/agro-platform/internal/testutil/postgres"
)

var ctx = context.Background()

func TestTaskRepo_Postgres_SaveAndGetByID(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	farmID := vo.NewID()
	tsk := task.New(farmID, "Полить рассаду")
	tsk.PullEvents()

	if err := repo.Save(ctx, tsk); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, err := repo.GetByID(ctx, tsk.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Title != "Полить рассаду" {
		t.Errorf("title: got %s", got.Title)
	}
	if got.Status != task.Todo {
		t.Errorf("status: got %s, want TODO", got.Status)
	}
}

func TestTaskRepo_Postgres_GetByID_NotFound(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	_, err := repo.GetByID(ctx, vo.NewID())
	if err != task.ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskRepo_Postgres_Upsert(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	tsk := task.New(vo.NewID(), "Старый")
	tsk.PullEvents()
	repo.Save(ctx, tsk)

	tsk.Title = "Новый"
	tsk.Start()
	repo.Save(ctx, tsk)

	got, _ := repo.GetByID(ctx, tsk.ID)
	if got.Title != "Новый" {
		t.Errorf("title after upsert: got %s", got.Title)
	}
	if got.Status != task.InProgress {
		t.Errorf("status after upsert: got %s", got.Status)
	}
}

func TestTaskRepo_Postgres_AssignmentRoundtrip(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	userID := vo.NewID()
	tsk := task.New(vo.NewID(), "С исполнителем")
	tsk.Assign(userID)
	tsk.PullEvents()
	repo.Save(ctx, tsk)

	got, _ := repo.GetByID(ctx, tsk.ID)
	if got.Assignment == nil {
		t.Fatal("assignment should be persisted")
	}
	if got.Assignment.UserID != userID {
		t.Errorf("userID: got %s, want %s", got.Assignment.UserID, userID)
	}
}

func TestTaskRepo_Postgres_CompletedAtRoundtrip(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	tsk := task.New(vo.NewID(), "Завершённая")
	tsk.Complete()
	tsk.PullEvents()
	repo.Save(ctx, tsk)

	got, _ := repo.GetByID(ctx, tsk.ID)
	if got.CompletedAt == nil {
		t.Fatal("CompletedAt should be persisted")
	}
}

func TestTaskRepo_Postgres_ListFiltersByFarm(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	farm1, farm2 := vo.NewID(), vo.NewID()
	for _, title := range []string{"A", "B"} {
		tsk := task.New(farm1, title)
		tsk.PullEvents()
		repo.Save(ctx, tsk)
	}
	tsk := task.New(farm2, "C")
	tsk.PullEvents()
	repo.Save(ctx, tsk)

	tasks, err := repo.List(ctx, repository.TaskFilter{FarmID: farm1})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks for farm1, got %d", len(tasks))
	}
}

func TestTaskRepo_Postgres_ListEmpty(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	tasks, err := repo.List(ctx, repository.TaskFilter{FarmID: vo.NewID()})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if tasks == nil {
		t.Error("expected empty slice, not nil")
	}
}

func TestTaskRepo_Postgres_OptionalFieldsRoundtrip(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	cycleID := vo.NewID()
	unitID := vo.NewID()
	desc := "Описание"

	tsk := task.New(vo.NewID(), "Полная задача")
	tsk.Description = &desc
	tsk.GrowingCycleID = &cycleID
	tsk.ProductionUnitID = &unitID
	tsk.Priority = task.Critical
	tsk.PullEvents()
	repo.Save(ctx, tsk)

	got, _ := repo.GetByID(ctx, tsk.ID)
	if got.Description == nil || *got.Description != desc {
		t.Error("description mismatch")
	}
	if got.GrowingCycleID == nil || *got.GrowingCycleID != cycleID {
		t.Error("growingCycleID mismatch")
	}
	if got.ProductionUnitID == nil || *got.ProductionUnitID != unitID {
		t.Error("productionUnitID mismatch")
	}
	if got.Priority != task.Critical {
		t.Errorf("priority: got %s, want CRITICAL", got.Priority)
	}
}

func TestTaskRepo_Postgres_Delete(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTaskRepository(db.Pool)

	tsk := task.New(vo.NewID(), "Удалить")
	tsk.PullEvents()
	repo.Save(ctx, tsk)
	repo.Delete(ctx, tsk.ID)

	_, err := repo.GetByID(ctx, tsk.ID)
	if err == nil {
		t.Fatal("expected error after delete")
	}
}
