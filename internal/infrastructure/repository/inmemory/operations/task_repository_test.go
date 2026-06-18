package operations_test

import (
	"context"
	"testing"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/task"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	inmemops "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/operations"
)

var ctx = context.Background()

func newTaskRepo() repository.TaskRepository {
	return inmemops.NewTaskRepository()
}

func makeTask(farmID vo.ID, title string) *task.Task {
	t := task.New(farmID, title)
	t.PullEvents()
	return t
}

func TestTaskRepo_SaveAndGetByID(t *testing.T) {
	repo := newTaskRepo()
	farmID := vo.NewID()
	tsk := makeTask(farmID, "Полить")

	if err := repo.Save(ctx, tsk); err != nil {
		t.Fatalf("save: %v", err)
	}

	got, err := repo.GetByID(ctx, tsk.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.ID != tsk.ID {
		t.Errorf("ID mismatch")
	}
	if got.Title != "Полить" {
		t.Errorf("title: got %s, want Полить", got.Title)
	}
}

func TestTaskRepo_GetByID_NotFound(t *testing.T) {
	repo := newTaskRepo()
	_, err := repo.GetByID(ctx, vo.NewID())
	if err == nil {
		t.Fatal("expected error for missing task")
	}
	if err != task.ErrTaskNotFound {
		t.Errorf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestTaskRepo_Save_Overwrites(t *testing.T) {
	repo := newTaskRepo()
	farmID := vo.NewID()
	tsk := makeTask(farmID, "Старый заголовок")
	repo.Save(ctx, tsk)

	tsk.Title = "Новый заголовок"
	repo.Save(ctx, tsk)

	got, _ := repo.GetByID(ctx, tsk.ID)
	if got.Title != "Новый заголовок" {
		t.Errorf("title after update: got %s, want Новый заголовок", got.Title)
	}
}

func TestTaskRepo_List_FiltersByFarmID(t *testing.T) {
	repo := newTaskRepo()
	farm1 := vo.NewID()
	farm2 := vo.NewID()

	repo.Save(ctx, makeTask(farm1, "A"))
	repo.Save(ctx, makeTask(farm1, "B"))
	repo.Save(ctx, makeTask(farm2, "C"))

	tasks, err := repo.List(ctx, repository.TaskFilter{FarmID: farm1})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks for farm1, got %d", len(tasks))
	}
}

func TestTaskRepo_List_FiltersByStatus(t *testing.T) {
	repo := newTaskRepo()
	farmID := vo.NewID()

	todo := makeTask(farmID, "todo task")
	done := makeTask(farmID, "done task")
	done.Complete()

	repo.Save(ctx, todo)
	repo.Save(ctx, done)

	doneStatus := task.Done
	tasks, _ := repo.List(ctx, repository.TaskFilter{FarmID: farmID, Status: &doneStatus})
	if len(tasks) != 1 {
		t.Errorf("expected 1 done task, got %d", len(tasks))
	}
	if tasks[0].Status != task.Done {
		t.Errorf("status: got %s, want DONE", tasks[0].Status)
	}
}

func TestTaskRepo_List_EmptyReturnsEmptySlice(t *testing.T) {
	repo := newTaskRepo()
	tasks, err := repo.List(ctx, repository.TaskFilter{FarmID: vo.NewID()})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if tasks == nil {
		t.Error("expected empty slice, got nil")
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	repo := newTaskRepo()
	farmID := vo.NewID()
	tsk := makeTask(farmID, "Удалить")
	repo.Save(ctx, tsk)

	repo.Delete(ctx, tsk.ID)

	_, err := repo.GetByID(ctx, tsk.ID)
	if err == nil {
		t.Fatal("expected error after delete")
	}
}
