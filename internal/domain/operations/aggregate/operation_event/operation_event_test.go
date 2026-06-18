package operationevent_test

import (
	"testing"

	operationevent "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

var farmID = vo.NewID()

// --- New ---

func TestNew_DefaultsAreCorrect(t *testing.T) {
	e := operationevent.New(farmID, operationevent.Irrigated)

	if e.ID.IsZero() {
		t.Error("expected non-zero ID")
	}
	if e.FarmID != farmID {
		t.Errorf("farmID: got %s, want %s", e.FarmID, farmID)
	}
	if e.Type != operationevent.Irrigated {
		t.Errorf("type: got %s, want IRRIGATED", e.Type)
	}
	if e.Timestamp.IsZero() {
		t.Error("Timestamp should not be zero")
	}
	if e.Payload == nil {
		t.Error("Payload should be initialized, not nil")
	}
}

func TestNew_EmitsOperationRecordedEvent(t *testing.T) {
	e := operationevent.New(farmID, operationevent.Irrigated)
	events := e.PullEvents()

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
	if events[0].EventType() != operationevent.EventOperationRecorded {
		t.Errorf("event type: got %s, want %s", events[0].EventType(), operationevent.EventOperationRecorded)
	}
}

// --- Payload ---

func TestNew_PayloadCanBePopulated(t *testing.T) {
	e := operationevent.New(farmID, operationevent.Fertilized)
	e.Payload["volumeLiters"] = 120.0
	e.Payload["method"] = "drip"

	if e.Payload["volumeLiters"] != 120.0 {
		t.Error("payload volumeLiters mismatch")
	}
	if e.Payload["method"] != "drip" {
		t.Error("payload method mismatch")
	}
}

// --- References ---

func TestNew_CanBeLinkedToProductionUnit(t *testing.T) {
	e := operationevent.New(farmID, operationevent.Irrigated)
	unitID := vo.NewID()
	e.ProductionUnitID = &unitID

	if e.ProductionUnitID == nil || *e.ProductionUnitID != unitID {
		t.Error("ProductionUnitID not set correctly")
	}
}

func TestNew_CanBeLinkedToGrowingCycle(t *testing.T) {
	e := operationevent.New(farmID, operationevent.Fertilized)
	cycleID := vo.NewID()
	e.GrowingCycleID = &cycleID

	if e.GrowingCycleID == nil || *e.GrowingCycleID != cycleID {
		t.Error("GrowingCycleID not set correctly")
	}
}

func TestNew_CanBeLinkedToPerformer(t *testing.T) {
	e := operationevent.New(farmID, operationevent.PHAdjusted)
	userID := vo.NewID()
	e.PerformedBy = &userID

	if e.PerformedBy == nil || *e.PerformedBy != userID {
		t.Error("PerformedBy not set correctly")
	}
}

// --- Operation types ---

func TestOperationTypes_AllAreValid(t *testing.T) {
	types := []operationevent.OperationType{
		operationevent.Irrigated,
		operationevent.Fertilized,
		operationevent.PHAdjusted,
		operationevent.Harvested,
		operationevent.Sprayed,
		operationevent.SensorAlert,
		operationevent.Transplanted,
	}

	for _, ot := range types {
		e := operationevent.New(farmID, ot)
		if e.Type != ot {
			t.Errorf("type mismatch: got %s, want %s", e.Type, ot)
		}
	}
}
