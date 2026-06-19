//go:build integration

package operations_test

import (
	"testing"
	"time"

	operationevent "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pgops "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/operations"
	dbtest "github.com/samurenkoroma/agro-platform/internal/testutil/postgres"
)

func TestOperationRepo_Postgres_SaveAndGetByID(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewOperationRepository(db.Pool)

	farmID := vo.NewID()
	e := operationevent.New(farmID, operationevent.Irrigated)
	e.PullEvents()

	if err := repo.Save(ctx, e); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, err := repo.GetByID(ctx, e.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Type != operationevent.Irrigated {
		t.Errorf("type: got %s, want IRRIGATED", got.Type)
	}
}

func TestOperationRepo_Postgres_JSONBPayloadRoundtrip(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewOperationRepository(db.Pool)

	e := operationevent.New(vo.NewID(), operationevent.Fertilized)
	e.Payload["amountGrams"] = 500.0
	e.Payload["productName"] = "NPK 20-20-20"
	e.Payload["targetEC"] = 2.8
	e.PullEvents()
	repo.Save(ctx, e)

	got, _ := repo.GetByID(ctx, e.ID)
	if got.Payload["amountGrams"] != 500.0 {
		t.Errorf("amountGrams: got %v", got.Payload["amountGrams"])
	}
	if got.Payload["productName"] != "NPK 20-20-20" {
		t.Errorf("productName: got %v", got.Payload["productName"])
	}
}

func TestOperationRepo_Postgres_IdempotentOnConflict(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewOperationRepository(db.Pool)

	e := operationevent.New(vo.NewID(), operationevent.Irrigated)
	e.PullEvents()

	if err := repo.Save(ctx, e); err != nil {
		t.Fatalf("first save: %v", err)
	}
	if err := repo.Save(ctx, e); err != nil {
		t.Fatalf("second save (ON CONFLICT DO NOTHING): %v", err)
	}
}

func TestOperationRepo_Postgres_ListFiltersByFarm(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewOperationRepository(db.Pool)

	farm1, farm2 := vo.NewID(), vo.NewID()

	for _, ot := range []operationevent.OperationType{operationevent.Irrigated, operationevent.Fertilized} {
		e := operationevent.New(farm1, ot)
		e.PullEvents()
		repo.Save(ctx, e)
	}
	e := operationevent.New(farm2, operationevent.Sprayed)
	e.PullEvents()
	repo.Save(ctx, e)

	ops, err := repo.List(ctx, repository.OperationFilter{FarmID: farm1})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(ops) != 2 {
		t.Errorf("expected 2 for farm1, got %d", len(ops))
	}
}

func TestOperationRepo_Postgres_ListEmpty(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewOperationRepository(db.Pool)

	ops, err := repo.List(ctx, repository.OperationFilter{FarmID: vo.NewID()})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if ops == nil {
		t.Error("expected empty slice, not nil")
	}
}

func TestOperationRepo_Postgres_ListOrderedByTimestampDesc(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewOperationRepository(db.Pool)

	farmID := vo.NewID()

	first := operationevent.New(farmID, operationevent.Irrigated)
	first.PullEvents()
	repo.Save(ctx, first)

	second := operationevent.New(farmID, operationevent.Fertilized)
	second.Timestamp = first.Timestamp.Add(time.Hour)
	second.PullEvents()
	repo.Save(ctx, second)

	ops, _ := repo.List(ctx, repository.OperationFilter{FarmID: farmID})
	if len(ops) < 2 {
		t.Fatalf("expected 2 ops, got %d", len(ops))
	}
	if ops[0].ID != second.ID {
		t.Error("expected most recent operation first")
	}
}

func TestOperationRepo_Postgres_RefsRoundtrip(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewOperationRepository(db.Pool)

	cycleID := vo.NewID()
	unitID := vo.NewID()
	userID := vo.NewID()

	e := operationevent.New(vo.NewID(), operationevent.PHAdjusted)
	e.GrowingCycleID = &cycleID
	e.ProductionUnitID = &unitID
	e.PerformedBy = &userID
	e.PullEvents()
	repo.Save(ctx, e)

	got, _ := repo.GetByID(ctx, e.ID)
	if got.GrowingCycleID == nil || *got.GrowingCycleID != cycleID {
		t.Error("GrowingCycleID mismatch")
	}
	if got.ProductionUnitID == nil || *got.ProductionUnitID != unitID {
		t.Error("ProductionUnitID mismatch")
	}
	if got.PerformedBy == nil || *got.PerformedBy != userID {
		t.Error("PerformedBy mismatch")
	}
}
