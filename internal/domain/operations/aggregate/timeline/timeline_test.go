package timeline_test

import (
	"testing"
	"time"

	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/timeline"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

var farmID = vo.NewID()

func newTimeline(t *testing.T) *timeline.Timeline {
	t.Helper()
	tl := timeline.New(farmID)
	tl.PullEvents()
	return tl
}

func newItem() timeline.Item {
	return timeline.Item{
		ID:          vo.NewID(),
		Source:      timeline.OperationSource,
		ReferenceID: vo.NewID(),
		Title:       "Полив",
		Timestamp:   time.Now(),
		Metadata:    vo.NewMetadata(),
	}
}

// --- New ---

func TestNew_DefaultsAreCorrect(t *testing.T) {
	tl := timeline.New(farmID)

	if tl.ID.IsZero() {
		t.Error("expected non-zero ID")
	}
	if tl.FarmID != farmID {
		t.Errorf("farmID: got %s, want %s", tl.FarmID, farmID)
	}
	if tl.Items == nil {
		t.Error("Items should be initialized, not nil")
	}
	if len(tl.Items) != 0 {
		t.Errorf("Items should be empty, got %d", len(tl.Items))
	}
}

func TestNew_EmitsTimelineCreatedEvent(t *testing.T) {
	tl := timeline.New(farmID)
	events := tl.PullEvents()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].EventType() != timeline.EventTimelineCreated {
		t.Errorf("event type: got %s, want %s", events[0].EventType(), timeline.EventTimelineCreated)
	}
}

// --- AddItem ---

func TestAddItem_AppendsItem(t *testing.T) {
	tl := newTimeline(t)
	item := newItem()
	tl.AddItem(item)

	if len(tl.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(tl.Items))
	}
	if tl.Items[0].ID != item.ID {
		t.Error("item ID mismatch")
	}
}

func TestAddItem_EmitsItemAddedEvent(t *testing.T) {
	tl := newTimeline(t)
	item := newItem()
	tl.AddItem(item)

	events := tl.PullEvents()
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].EventType() != timeline.EventTimelineItemAdded {
		t.Errorf("event type: got %s, want %s", events[0].EventType(), timeline.EventTimelineItemAdded)
	}
}

func TestAddItem_MultipleItemsPreserveOrder(t *testing.T) {
	tl := newTimeline(t)

	first := newItem()
	first.Title = "first"
	second := newItem()
	second.Title = "second"
	third := newItem()
	third.Title = "third"

	tl.AddItem(first)
	tl.AddItem(second)
	tl.AddItem(third)

	if len(tl.Items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(tl.Items))
	}
	if tl.Items[0].Title != "first" || tl.Items[1].Title != "second" || tl.Items[2].Title != "third" {
		t.Error("items order is wrong")
	}
}

func TestAddItem_UpdatesUpdatedAt(t *testing.T) {
	tl := newTimeline(t)
	before := tl.UpdatedAt
	time.Sleep(time.Millisecond)
	tl.AddItem(newItem())

	if !tl.UpdatedAt.After(before) {
		t.Error("UpdatedAt should be updated after AddItem")
	}
}

// --- GrowingCycleID ---

func TestTimeline_CanBeLinkedToGrowingCycle(t *testing.T) {
	tl := newTimeline(t)
	cycleID := vo.NewID()
	tl.GrowingCycleID = &cycleID

	if tl.GrowingCycleID == nil {
		t.Fatal("GrowingCycleID should be set")
	}
	if *tl.GrowingCycleID != cycleID {
		t.Error("GrowingCycleID mismatch")
	}
}

// --- Source types ---

func TestItem_SourceTypes(t *testing.T) {
	sources := []timeline.Source{
		timeline.TaskSource,
		timeline.OperationSource,
		timeline.HarvestSource,
	}
	tl := newTimeline(t)

	for _, src := range sources {
		item := newItem()
		item.Source = src
		tl.AddItem(item)
	}

	if len(tl.Items) != len(sources) {
		t.Errorf("expected %d items, got %d", len(sources), len(tl.Items))
	}
}
