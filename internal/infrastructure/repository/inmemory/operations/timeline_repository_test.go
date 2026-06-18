package operations_test

import (
	"testing"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/timeline"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	inmemops "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/operations"
)

func TestTimelineRepo_SaveAndGetByID(t *testing.T) {
	repo := inmemops.NewTimelineRepository()
	farmID := vo.NewID()
	tl := timeline.New(farmID)

	if err := repo.Save(ctx, tl); err != nil {
		t.Fatalf("save: %v", err)
	}

	got, err := repo.GetByID(ctx, tl.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.ID != tl.ID {
		t.Error("ID mismatch")
	}
}

func TestTimelineRepo_GetByOwner_WithCycleID(t *testing.T) {
	repo := inmemops.NewTimelineRepository()
	farmID := vo.NewID()
	cycleID := vo.NewID()

	tl := timeline.New(farmID)
	tl.GrowingCycleID = &cycleID
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

func TestTimelineRepo_GetByOwner_NilCycle(t *testing.T) {
	repo := inmemops.NewTimelineRepository()
	farmID := vo.NewID()

	tl := timeline.New(farmID)
	repo.Save(ctx, tl)

	got, _ := repo.GetByOwner(ctx, farmID, nil)
	if got == nil {
		t.Fatal("expected timeline without cycle, got nil")
	}
}

func TestTimelineRepo_GetByOwner_NotFound(t *testing.T) {
	repo := inmemops.NewTimelineRepository()
	got, err := repo.GetByOwner(ctx, vo.NewID(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != nil {
		t.Error("expected nil for missing timeline")
	}
}

func TestTimelineRepo_Save_UpdatesExisting(t *testing.T) {
	repo := inmemops.NewTimelineRepository()
	farmID := vo.NewID()
	tl := timeline.New(farmID)
	repo.Save(ctx, tl)

	tl.AddItem(timeline.Item{
		ID:          vo.NewID(),
		Source:      timeline.OperationSource,
		ReferenceID: vo.NewID(),
		Title:       "item",
		Metadata:    vo.NewMetadata(),
	})
	repo.Save(ctx, tl)

	got, _ := repo.GetByID(ctx, tl.ID)
	if len(got.Items) != 1 {
		t.Errorf("expected 1 item after update, got %d", len(got.Items))
	}
}

func TestTimelineRepo_GetByOwner_IsolatesByFarm(t *testing.T) {
	repo := inmemops.NewTimelineRepository()
	farm1 := vo.NewID()
	farm2 := vo.NewID()

	repo.Save(ctx, timeline.New(farm1))
	repo.Save(ctx, timeline.New(farm2))

	got1, _ := repo.GetByOwner(ctx, farm1, nil)
	got2, _ := repo.GetByOwner(ctx, farm2, nil)

	if got1.FarmID != farm1 {
		t.Error("farm1 timeline has wrong FarmID")
	}
	if got2.FarmID != farm2 {
		t.Error("farm2 timeline has wrong FarmID")
	}
}
