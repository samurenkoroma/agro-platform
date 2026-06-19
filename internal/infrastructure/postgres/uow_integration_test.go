//go:build integration

package postgres_test

import (
	"context"
	"errors"
	"testing"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/task"
	opsrepo "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/infrastructure/messaging/inmemory"
	pgUow "github.com/samurenkoroma/agro-platform/internal/infrastructure/postgres"
	pgops "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/operations"
	"github.com/samurenkoroma/agro-platform/internal/shared/repository"
	dbtest "github.com/samurenkoroma/agro-platform/internal/testutil/postgres"
)

type opsProvider struct{ db uow.DB }

func (p *opsProvider) ProviderName() string          { return "operations" }
func (p *opsProvider) Tasks() opsrepo.TaskRepository { return pgops.NewTaskRepository(p.db) }
func (p *opsProvider) Timelines() opsrepo.TimeLineRepository {
	return pgops.NewTimelineRepository(p.db)
}
func (p *opsProvider) Operations() opsrepo.OperationRepository {
	return pgops.NewOperationRepository(p.db)
}

func buildOpsProvider(db uow.DB) repository.RepositoryProvider {
	return &opsProvider{db: db}
}

// --- Atomicity ---

func TestUoW_Postgres_CommitsOnSuccess(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	ctx := context.Background()
	bus := inmemory.NewInMemoryEventBus()
	u := pgUow.NewUnitOfWork(db.Pool, bus)

	farmID := vo.NewID()

	_, err := u.Execute(ctx, buildOpsProvider, func(p repository.RepositoryProvider, exec uow.Execution) (any, error) {
		ops := p.(opsrepo.OperationsProvider)
		tsk := task.New(farmID, "Committed task")
		if err := ops.Tasks().Save(ctx, tsk); err != nil {
			return nil, err
		}
		exec.RegisterAggregate(tsk)
		return nil, nil
	})
	if err != nil {
		t.Fatalf("execute: %v", err)
	}

	// Проверяем напрямую через новый репозиторий на том же pool —
	// если транзакция не закоммитилась, строки не будет.
	repo := pgops.NewTaskRepository(db.Pool)
	tasks, err := repo.List(ctx, opsrepo.TaskFilter{FarmID: farmID})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected 1 committed task, got %d", len(tasks))
	}
}

func TestUoW_Postgres_RollsBackOnError(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	ctx := context.Background()
	bus := inmemory.NewInMemoryEventBus()
	u := pgUow.NewUnitOfWork(db.Pool, bus)

	farmID := vo.NewID()
	boom := errors.New("boom")

	_, err := u.Execute(ctx, buildOpsProvider, func(p repository.RepositoryProvider, exec uow.Execution) (any, error) {
		ops := p.(opsrepo.OperationsProvider)
		tsk := task.New(farmID, "Should be rolled back")
		if err := ops.Tasks().Save(ctx, tsk); err != nil {
			return nil, err
		}
		// Искусственная ошибка после Save — должна откатить всё, включая Save выше.
		return nil, boom
	})
	if err == nil {
		t.Fatal("expected error to propagate")
	}

	repo := pgops.NewTaskRepository(db.Pool)
	tasks, _ := repo.List(ctx, opsrepo.TaskFilter{FarmID: farmID})
	if len(tasks) != 0 {
		t.Fatalf("expected 0 tasks after rollback, got %d", len(tasks))
	}
}

func TestUoW_Postgres_MultipleAggregatesAreAtomic(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	ctx := context.Background()
	bus := inmemory.NewInMemoryEventBus()
	u := pgUow.NewUnitOfWork(db.Pool, bus)

	farmID := vo.NewID()
	boom := errors.New("boom after second save")

	_, err := u.Execute(ctx, buildOpsProvider, func(p repository.RepositoryProvider, exec uow.Execution) (any, error) {
		ops := p.(opsrepo.OperationsProvider)

		t1 := task.New(farmID, "First")
		if err := ops.Tasks().Save(ctx, t1); err != nil {
			return nil, err
		}
		t2 := task.New(farmID, "Second")
		if err := ops.Tasks().Save(ctx, t2); err != nil {
			return nil, err
		}
		// Ошибка после ДВУХ успешных Save — оба должны откатиться.
		return nil, boom
	})
	if err == nil {
		t.Fatal("expected error")
	}

	repo := pgops.NewTaskRepository(db.Pool)
	tasks, _ := repo.List(ctx, opsrepo.TaskFilter{FarmID: farmID})
	if len(tasks) != 0 {
		t.Fatalf("expected both saves rolled back, got %d tasks", len(tasks))
	}
}

// --- Per-call isolation (регрессия для старого бага: одно состояние на весь uow) ---

func TestUoW_Postgres_SameInstanceHandlesMultipleSequentialExecutes(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	ctx := context.Background()
	bus := inmemory.NewInMemoryEventBus()
	u := pgUow.NewUnitOfWork(db.Pool, bus)

	farmID := vo.NewID()

	// Раньше committed/rolledBack/aggregates жили на самом unitOfWork —
	// второй Execute на том же инстансе падал с ErrAlreadyCommitted
	// или повторно диспатчил события первого вызова. Теперь это per-call
	// состояние внутри exec, так что несколько Execute подряд на одном uow
	// (как и происходит в реальном приложении — один uow на всё) должны
	// работать независимо.
	for i := 0; i < 3; i++ {
		_, err := u.Execute(ctx, buildOpsProvider, func(p repository.RepositoryProvider, exec uow.Execution) (any, error) {
			ops := p.(opsrepo.OperationsProvider)
			tsk := task.New(farmID, "Sequential task")
			if err := ops.Tasks().Save(ctx, tsk); err != nil {
				return nil, err
			}
			exec.RegisterAggregate(tsk)
			return nil, nil
		})
		if err != nil {
			t.Fatalf("execute #%d: %v", i+1, err)
		}
	}

	repo := pgops.NewTaskRepository(db.Pool)
	tasks, _ := repo.List(ctx, opsrepo.TaskFilter{FarmID: farmID})
	if len(tasks) != 3 {
		t.Fatalf("expected 3 tasks after 3 sequential Executes on same uow, got %d", len(tasks))
	}
}
