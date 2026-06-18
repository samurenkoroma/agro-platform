package operations_test

import (
	"testing"

	operationevent "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	inmemops "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/operations"
)

func TestOperationRepo_SaveAndGetByID(t *testing.T) {
	repo := inmemops.NewOperationRepository()
	farmID := vo.NewID()
	e := operationevent.New(farmID, operationevent.Irrigated)

	if err := repo.Save(ctx, e); err != nil {
		t.Fatalf("save: %v", err)
	}

	got, err := repo.GetByID(ctx, e.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.ID != e.ID {
		t.Error("ID mismatch")
	}
	if got.Type != operationevent.Irrigated {
		t.Errorf("type: got %s, want IRRIGATED", got.Type)
	}
}

func TestOperationRepo_GetByID_NotFound(t *testing.T) {
	repo := inmemops.NewOperationRepository()
	_, err := repo.GetByID(ctx, vo.NewID())
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
	if err != operationevent.ErrOperationNotFound {
		t.Errorf("expected ErrOperationNotFound, got %v", err)
	}
}

func TestOperationRepo_List_FiltersByFarmID(t *testing.T) {
	repo := inmemops.NewOperationRepository()
	farm1 := vo.NewID()
	farm2 := vo.NewID()

	repo.Save(ctx, operationevent.New(farm1, operationevent.Irrigated))
	repo.Save(ctx, operationevent.New(farm1, operationevent.Fertilized))
	repo.Save(ctx, operationevent.New(farm2, operationevent.Sprayed))

	ops, err := repo.List(ctx, repository.OperationFilter{FarmID: farm1})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(ops) != 2 {
		t.Errorf("expected 2 operations for farm1, got %d", len(ops))
	}
}

func TestOperationRepo_List_FiltersByType(t *testing.T) {
	repo := inmemops.NewOperationRepository()
	farmID := vo.NewID()

	repo.Save(ctx, operationevent.New(farmID, operationevent.Irrigated))
	repo.Save(ctx, operationevent.New(farmID, operationevent.Fertilized))
	repo.Save(ctx, operationevent.New(farmID, operationevent.Irrigated))

	opType := operationevent.Irrigated
	ops, _ := repo.List(ctx, repository.OperationFilter{FarmID: farmID, Type: &opType})
	if len(ops) != 2 {
		t.Errorf("expected 2 IRRIGATED operations, got %d", len(ops))
	}
}

func TestOperationRepo_List_EmptyReturnsEmptySlice(t *testing.T) {
	repo := inmemops.NewOperationRepository()
	ops, err := repo.List(ctx, repository.OperationFilter{FarmID: vo.NewID()})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if ops == nil {
		t.Error("expected empty slice, got nil")
	}
}
