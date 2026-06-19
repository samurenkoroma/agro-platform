//go:build integration

package operations_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/samurenkoroma/agro-platform/internal/testutil/e2e"
)

func newClient(t *testing.T) *e2e.Client {
	t.Helper()
	c := e2e.NewClient(t, "account", "operations")
	c.SetupOrg(t, "farmer@example.com", "farmer", "secret123", "Тестовая ферма")
	return c
}

type idData struct {
	ID string `json:"id"`
}

type taskDTO struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// --- Lifecycle ---

func TestE2E_Task_Create(t *testing.T) {
	c := newClient(t)

	env := c.Command("operations.create_task", map[string]any{
		"title": "Полить рассаду", "priority": "HIGH",
	})
	if !env.Success {
		t.Fatalf("create_task: %+v", env.Error)
	}
	var d idData
	json.Unmarshal(env.Data, &d)
	if d.ID == "" {
		t.Fatal("expected non-empty task ID")
	}
}

func TestE2E_Task_GetAfterCreate(t *testing.T) {
	c := newClient(t)

	createEnv := c.Command("operations.create_task", map[string]any{"title": "Задача"})
	var created idData
	json.Unmarshal(createEnv.Data, &created)

	getEnv := c.Query("operations.get_task", map[string]any{"id": created.ID})
	if !getEnv.Success {
		t.Fatalf("get_task: %+v", getEnv.Error)
	}
	var task taskDTO
	json.Unmarshal(getEnv.Data, &task)
	if task.Status != "TODO" {
		t.Errorf("status: got %s, want TODO", task.Status)
	}
	if task.Title != "Задача" {
		t.Errorf("title: got %s", task.Title)
	}
}

func TestE2E_Task_StartChangesStatus(t *testing.T) {
	c := newClient(t)

	env := c.Command("operations.create_task", map[string]any{"title": "Задача"})
	var created idData
	json.Unmarshal(env.Data, &created)

	if r := c.Command("operations.start_task", map[string]any{"taskId": created.ID}); !r.Success {
		t.Fatalf("start_task: %+v", r.Error)
	}

	getEnv := c.Query("operations.get_task", map[string]any{"id": created.ID})
	var task taskDTO
	json.Unmarshal(getEnv.Data, &task)
	if task.Status != "IN_PROGRESS" {
		t.Errorf("status: got %s, want IN_PROGRESS", task.Status)
	}
}

func TestE2E_Task_CompleteChangesStatus(t *testing.T) {
	c := newClient(t)

	env := c.Command("operations.create_task", map[string]any{"title": "Задача"})
	var created idData
	json.Unmarshal(env.Data, &created)

	c.Command("operations.complete_task", map[string]any{"taskId": created.ID})

	var task taskDTO
	json.Unmarshal(c.Query("operations.get_task", map[string]any{"id": created.ID}).Data, &task)
	if task.Status != "DONE" {
		t.Errorf("status: got %s, want DONE", task.Status)
	}
}

func TestE2E_Task_CancelChangesStatus(t *testing.T) {
	c := newClient(t)

	env := c.Command("operations.create_task", map[string]any{"title": "Задача"})
	var created idData
	json.Unmarshal(env.Data, &created)

	c.Command("operations.cancel_task", map[string]any{"taskId": created.ID})

	var task taskDTO
	json.Unmarshal(c.Query("operations.get_task", map[string]any{"id": created.ID}).Data, &task)
	if task.Status != "CANCELLED" {
		t.Errorf("status: got %s, want CANCELLED", task.Status)
	}
}

func TestE2E_Task_FullLifecycle(t *testing.T) {
	c := newClient(t)

	env := c.Command("operations.create_task", map[string]any{"title": "Полный цикл"})
	var created idData
	json.Unmarshal(env.Data, &created)

	check := func(wantStatus string) {
		t.Helper()
		var task taskDTO
		json.Unmarshal(c.Query("operations.get_task", map[string]any{"id": created.ID}).Data, &task)
		if task.Status != wantStatus {
			t.Errorf("status: got %s, want %s", task.Status, wantStatus)
		}
	}

	check("TODO")
	c.Command("operations.start_task", map[string]any{"taskId": created.ID})
	check("IN_PROGRESS")
	c.Command("operations.complete_task", map[string]any{"taskId": created.ID})
	check("DONE")
}

func TestE2E_Task_Assign(t *testing.T) {
	c := newClient(t)

	env := c.Command("operations.create_task", map[string]any{"title": "Назначаемая"})
	var created idData
	json.Unmarshal(env.Data, &created)

	assignEnv := c.Command("operations.assign_task", map[string]any{
		"taskId": created.ID,
		"userId": "11111111-1111-1111-1111-111111111111",
	})
	if !assignEnv.Success {
		t.Fatalf("assign_task: %+v", assignEnv.Error)
	}
}

func TestE2E_Task_List(t *testing.T) {
	c := newClient(t)

	for _, title := range []string{"Задача 1", "Задача 2", "Задача 3"} {
		c.Command("operations.create_task", map[string]any{"title": title})
	}

	listEnv := c.Query("operations.list_tasks", map[string]any{})
	if !listEnv.Success {
		t.Fatalf("list_tasks: %+v", listEnv.Error)
	}
	var tasks []taskDTO
	json.Unmarshal(listEnv.Data, &tasks)
	if len(tasks) != 3 {
		t.Errorf("expected 3 tasks, got %d", len(tasks))
	}
}

// --- Security ---

func TestE2E_Task_Unauthorized_Returns401(t *testing.T) {
	c := e2e.NewClient(t, "account", "operations")

	status, env := c.PostWithoutAuth("/api/commands", map[string]any{
		"command": "operations.create_task",
		"data":    map[string]any{"title": "test"},
	})
	if status != http.StatusUnauthorized {
		t.Errorf("status: got %d, want 401", status)
	}
	if env.Success {
		t.Error("expected success=false for unauthenticated request")
	}
}

// --- Isolation ---

func TestE2E_Task_IsolatedBetweenOrganizations(t *testing.T) {
	c1 := e2e.NewClient(t, "account", "operations")
	c1.SetupOrg(t, "farm1@example.com", "farm1", "secret123", "Ферма 1")
	c1.Command("operations.create_task", map[string]any{"title": "Задача фермы 1"})

	c2 := e2e.NewClient(t, "account", "operations")
	c2.SetupOrg(t, "farm2@example.com", "farm2", "secret123", "Ферма 2")

	var tasks1, tasks2 []taskDTO
	json.Unmarshal(c1.Query("operations.list_tasks", map[string]any{}).Data, &tasks1)
	json.Unmarshal(c2.Query("operations.list_tasks", map[string]any{}).Data, &tasks2)

	if len(tasks1) != 1 {
		t.Errorf("farm1: expected 1 task, got %d", len(tasks1))
	}
	if len(tasks2) != 0 {
		t.Errorf("farm2: expected 0 tasks, got %d", len(tasks2))
	}
}
