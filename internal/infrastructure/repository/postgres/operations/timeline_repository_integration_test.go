//go:build integration

package operations_test

import (
	"testing"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/timeline"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	pgops "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/postgres/operations"
	dbtest "github.com/samurenkoroma/agro-platform/internal/testutil/postgres"
)

func TestTimelineRepo_Postgres_SaveAndGetByID(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTimelineRepository(db.Pool)

	farmID := vo.NewID()
	tl := timeline.New(farmID)
	tl.PullEvents()

	if err := repo.Save(ctx, tl); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, err := repo.GetByID(ctx, tl.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.FarmID != farmID {
		t.Errorf("farmID mismatch")
	}
}

func TestTimelineRepo_Postgres_ItemsJSONBRoundtrip(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTimelineRepository(db.Pool)

	tl := timeline.New(vo.NewID())
	tl.PullEvents()
	tl.AddItem(timeline.Item{
		ID: vo.NewID(), Source: timeline.OperationSource,
		ReferenceID: vo.NewID(), Title: "Полив", Metadata: vo.NewMetadata(),
	})
	tl.AddItem(timeline.Item{
		ID: vo.NewID(), Source: timeline.TaskSource,
		ReferenceID: vo.NewID(), Title: "Удобрение", Metadata: vo.NewMetadata(),
	})

	repo.Save(ctx, tl)

	got, _ := repo.GetByID(ctx, tl.ID)
	if len(got.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(got.Items))
	}
	if got.Items[0].Title != "Полив" {
		t.Errorf("item[0] title: got %s", got.Items[0].Title)
	}
	if got.Items[1].Source != timeline.TaskSource {
		t.Errorf("item[1] source: got %s", got.Items[1].Source)
	}
}

func TestTimelineRepo_Postgres_GetByOwner_WithCycle(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTimelineRepository(db.Pool)

	farmID := vo.NewID()
	cycleID := vo.NewID()

	tl := timeline.New(farmID)
	tl.GrowingCycleID = &cycleID
	tl.PullEvents()
	repo.Save(ctx, tl)

	got, err := repo.GetByOwner(ctx, farmID, &cycleID)
	if err != nil {
		t.Fatalf("get by owner: %v", err)
	}
	if got == nil {
		t.Fatal("expected timeline, got nil")
	}
	if got.ID != tl.ID {
		t.Error("ID mismatch")
	}
}

func TestTimelineRepo_Postgres_GetByOwner_WithoutCycle(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTimelineRepository(db.Pool)

	farmID := vo.NewID()
	tl := timeline.New(farmID)
	tl.PullEvents()
	repo.Save(ctx, tl)

	got, _ := repo.GetByOwner(ctx, farmID, nil)
	if got == nil {
		t.Fatal("expected timeline without cycle")
	}
}

func TestTimelineRepo_Postgres_GetByOwner_NotFound(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTimelineRepository(db.Pool)

	got, err := repo.GetByOwner(ctx, vo.NewID(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != nil {
		t.Error("expected nil for missing timeline")
	}
}

func TestTimelineRepo_Postgres_UniqueOwnerConstraint(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTimelineRepository(db.Pool)

	farmID := vo.NewID()
	cycleID := vo.NewID()

	first := timeline.New(farmID)
	first.GrowingCycleID = &cycleID
	first.PullEvents()
	if err := repo.Save(ctx, first); err != nil {
		t.Fatalf("save first: %v", err)
	}

	// Второй таймлайн с тем же (farm_id, growing_cycle_id) — уникальный индекс.
	second := timeline.New(farmID)
	second.GrowingCycleID = &cycleID
	second.PullEvents()
	if err := repo.Save(ctx, second); err == nil {
		t.Error("expected unique constraint violation")
	}
}

func TestTimelineRepo_Postgres_UpdateAppendedItems(t *testing.T) {
	db := dbtest.NewTestDB(t, "operations")
	repo := pgops.NewTimelineRepository(db.Pool)

	tl := timeline.New(vo.NewID())
	tl.PullEvents()
	repo.Save(ctx, tl)

	tl.AddItem(timeline.Item{
		ID: vo.NewID(), Source: timeline.OperationSource,
		ReferenceID: vo.NewID(), Title: "Добавлено позже", Metadata: vo.NewMetadata(),
	})
	repo.Save(ctx, tl)

	got, _ := repo.GetByID(ctx, tl.ID)
	if len(got.Items) != 1 {
		t.Errorf("expected 1 item after update, got %d", len(got.Items))
	}
}
