package operations

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	operationevent "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/task"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type taskRepository struct{ db uow.DB }

func NewTaskRepository(db uow.DB) repository.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Save(ctx context.Context, t *task.Task) error {
	var assignedTo *string
	if t.Assignment != nil {
		s := string(t.Assignment.UserID)
		assignedTo = &s
	}
	_, err := r.db.Exec(ctx, `
		INSERT INTO operations_tasks (
			id, farm_id, title, description, operation_type,
			production_unit_id, growing_cycle_id,
			assigned_to, status, priority,
			due_date, completed_at,
			created_at, updated_at
		) VALUES (
			$1,$2,$3,$4,$5,
			$6,$7,
			$8,$9,$10,
			$11,$12,
			$13,$14
		)
		ON CONFLICT (id) DO UPDATE SET
			title             = EXCLUDED.title,
			description       = EXCLUDED.description,
			operation_type    = EXCLUDED.operation_type,
			production_unit_id = EXCLUDED.production_unit_id,
			growing_cycle_id  = EXCLUDED.growing_cycle_id,
			assigned_to       = EXCLUDED.assigned_to,
			status            = EXCLUDED.status,
			priority          = EXCLUDED.priority,
			due_date          = EXCLUDED.due_date,
			completed_at      = EXCLUDED.completed_at,
			updated_at        = EXCLUDED.updated_at
	`,
		t.ID, t.FarmID, t.Title, t.Description, t.OperationType,
		t.ProductionUnitID, t.GrowingCycleID,
		assignedTo, t.Status, t.Priority,
		t.DueDate, t.CompletedAt,
		t.CreatedAt, t.UpdatedAt,
	)
	return err
}

func (r *taskRepository) GetByID(ctx context.Context, id vo.ID) (*task.Task, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, farm_id, title, description, operation_type,
		       production_unit_id, growing_cycle_id,
		       assigned_to, status, priority,
		       due_date, completed_at, created_at, updated_at
		FROM operations_tasks WHERE id = $1`, id)
	return scanTask(row)
}

func (r *taskRepository) List(ctx context.Context, filter repository.TaskFilter) ([]*task.Task, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, farm_id, title, description, operation_type,
		       production_unit_id, growing_cycle_id,
		       assigned_to, status, priority,
		       due_date, completed_at, created_at, updated_at
		FROM operations_tasks
		WHERE farm_id = $1
		ORDER BY created_at DESC`, filter.FarmID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*task.Task
	for rows.Next() {
		t, err := scanTask(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (r *taskRepository) Delete(ctx context.Context, id vo.ID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM operations_tasks WHERE id = $1`, id)
	return err
}

type scanner interface{ Scan(dest ...any) error }

func scanTask(s scanner) (*task.Task, error) {
	var t task.Task
	var assignedTo *string
	var opType *string
	err := s.Scan(
		&t.ID, &t.FarmID, &t.Title, &t.Description, &opType,
		&t.ProductionUnitID, &t.GrowingCycleID,
		&assignedTo, &t.Status, &t.Priority,
		&t.DueDate, &t.CompletedAt, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, task.ErrTaskNotFound
		}
		return nil, err
	}
	if assignedTo != nil {
		t.Assignment = &task.Assignment{UserID: vo.ID(*assignedTo)}
	}
	if opType != nil {
		ot := operationevent.OperationType(*opType)
		t.OperationType = &ot
	}
	return &t, nil
}

var _ repository.TaskRepository = (*taskRepository)(nil)
