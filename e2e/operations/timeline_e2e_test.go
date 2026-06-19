//go:build integration

package operations_test

import (
	"encoding/json"
	"testing"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	"github.com/samurenkoroma/agro-platform/internal/testutil/e2e"
)

type operationDTO struct {
	ID      string         `json:"id"`
	Type    string         `json:"type"`
	Payload map[string]any `json:"payload"`
}

type timelineDTO struct {
	ID    string `json:"id"`
	Items []struct {
		Title string `json:"title"`
	} `json:"items"`
}

// --- record_operation ---

func TestE2E_RecordOperation_ReturnsID(t *testing.T) {
	c := e2e.NewClient(t, "account", "operations")
	c.SetupOrg(t, "op1@example.com", "op1", "secret123", "Ферма")

	env := c.Command("operations.record_operation", map[string]any{
		"type": "IRRIGATED",
	})
	if !env.Success {
		t.Fatalf("record_operation: %+v", env.Error)
	}
	var d idData
	json.Unmarshal(env.Data, &d)
	if d.ID == "" {
		t.Fatal("expected non-empty operation ID")
	}
}

func TestE2E_RecordOperation_AppearsInList(t *testing.T) {
	c := e2e.NewClient(t, "account", "operations")
	c.SetupOrg(t, "op2@example.com", "op2", "secret123", "Ферма")

	cycleID := vo.NewID().String()
	c.Command("operations.record_operation", map[string]any{
		"type": "IRRIGATED", "growingCycleId": cycleID,
	})

	listEnv := c.Query("operations.list_operations", map[string]any{"growingCycleId": cycleID})
	if !listEnv.Success {
		t.Fatalf("list_operations: %+v", listEnv.Error)
	}
	var ops []operationDTO
	json.Unmarshal(listEnv.Data, &ops)
	if len(ops) != 1 {
		t.Fatalf("expected 1 operation, got %d", len(ops))
	}
	if ops[0].Type != "IRRIGATED" {
		t.Errorf("type: got %s", ops[0].Type)
	}
}

func TestE2E_RecordOperation_PayloadRoundtrip(t *testing.T) {
	c := e2e.NewClient(t, "account", "operations")
	c.SetupOrg(t, "op3@example.com", "op3", "secret123", "Ферма")

	cycleID := vo.NewID().String()
	c.Command("operations.record_operation", map[string]any{
		"type":           "FERTILIZED",
		"growingCycleId": cycleID,
		"payload": map[string]any{
			"productName": "NPK 20-20-20",
			"amountGrams": 500,
		},
	})

	listEnv := c.Query("operations.list_operations", map[string]any{"growingCycleId": cycleID})
	var ops []operationDTO
	json.Unmarshal(listEnv.Data, &ops)

	if ops[0].Payload["productName"] != "NPK 20-20-20" {
		t.Errorf("productName: got %v", ops[0].Payload["productName"])
	}
}

// --- Timeline ---

func TestE2E_RecordOperation_AutoCreatesTimeline(t *testing.T) {
	c := e2e.NewClient(t, "account", "operations")
	c.SetupOrg(t, "op4@example.com", "op4", "secret123", "Ферма")

	cycleID := vo.NewID().String()
	c.Command("operations.record_operation", map[string]any{
		"type": "IRRIGATED", "growingCycleId": cycleID,
	})

	tlEnv := c.Query("operations.get_timeline", map[string]any{"growingCycleId": cycleID})
	if !tlEnv.Success {
		t.Fatalf("get_timeline: %+v", tlEnv.Error)
	}
	var tl timelineDTO
	json.Unmarshal(tlEnv.Data, &tl)
	if len(tl.Items) != 1 {
		t.Fatalf("expected 1 timeline item, got %d", len(tl.Items))
	}
	if tl.Items[0].Title != "IRRIGATED" {
		t.Errorf("item title: got %s", tl.Items[0].Title)
	}
}

func TestE2E_RecordOperation_AccumulatesInTimeline(t *testing.T) {
	c := e2e.NewClient(t, "account", "operations")
	c.SetupOrg(t, "op5@example.com", "op5", "secret123", "Ферма")

	cycleID := vo.NewID().String()
	for _, opType := range []string{"IRRIGATED", "FERTILIZED", "PH_ADJUSTED"} {
		c.Command("operations.record_operation", map[string]any{
			"type": opType, "growingCycleId": cycleID,
		})
	}

	var tl timelineDTO
	json.Unmarshal(c.Query("operations.get_timeline", map[string]any{"growingCycleId": cycleID}).Data, &tl)
	if len(tl.Items) != 3 {
		t.Fatalf("expected 3 timeline items, got %d", len(tl.Items))
	}
}

func TestE2E_RecordOperation_FarmWideTimeline(t *testing.T) {
	c := e2e.NewClient(t, "account", "operations")
	c.SetupOrg(t, "op6@example.com", "op6", "secret123", "Ферма")

	// Без growingCycleId — попадает в общий таймлайн фермы
	c.Command("operations.record_operation", map[string]any{"type": "INSPECTED"})

	tlEnv := c.Query("operations.get_timeline", map[string]any{})
	if !tlEnv.Success {
		t.Fatalf("get_timeline: %+v", tlEnv.Error)
	}
	var tl timelineDTO
	json.Unmarshal(tlEnv.Data, &tl)
	if len(tl.Items) == 0 {
		t.Error("expected at least 1 item in farm-wide timeline")
	}
}

func TestE2E_Operations_IsolatedBetweenOrgs(t *testing.T) {
	c1 := e2e.NewClient(t, "account", "operations")
	c1.SetupOrg(t, "iso1@example.com", "iso1", "secret123", "Ферма 1")
	c1.Command("operations.record_operation", map[string]any{"type": "IRRIGATED"})

	c2 := e2e.NewClient(t, "account", "operations")
	c2.SetupOrg(t, "iso2@example.com", "iso2", "secret123", "Ферма 2")

	var ops1, ops2 []operationDTO
	json.Unmarshal(c1.Query("operations.list_operations", map[string]any{}).Data, &ops1)
	json.Unmarshal(c2.Query("operations.list_operations", map[string]any{}).Data, &ops2)

	if len(ops1) != 1 {
		t.Errorf("org1: expected 1 operation, got %d", len(ops1))
	}
	if len(ops2) != 0 {
		t.Errorf("org2: expected 0 operations, got %d", len(ops2))
	}
}
