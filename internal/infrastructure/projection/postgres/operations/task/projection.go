package task

import (
	"context"

	taskquery "github.com/samurenkoroma/agro-platform/internal/application/queries/operations/task"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct{ db uow.DB }

func New(db uow.DB) taskquery.Projection { return &projection{db: db} }

func (p *projection) Get(ctx context.Context, id vo.ID) (*taskquery.TaskDTO, error) {
	row := p.db.QueryRow(ctx, `
		SELECT id, title, description, operation_type,
		       production_unit_id, growing_cycle_id,
		       assigned_to, status, priority,
		       due_date, completed_at, created_at
		FROM operations_tasks WHERE id = $1`, id)
	return scanTaskDTO(row)
}

func (p *projection) List(ctx context.Context, farmID vo.ID, cycleID *vo.ID) ([]*taskquery.TaskDTO, error) {
	var rows interface {
		Next() bool
		Scan(...any) error
		Close()
	}
	var err error
	if cycleID != nil {
		rows, err = p.db.Query(ctx, `
			SELECT id, title, description, operation_type,
			       production_unit_id, growing_cycle_id,
			       assigned_to, status, priority,
			       due_date, completed_at, created_at
			FROM operations_tasks
			WHERE farm_id = $1 AND growing_cycle_id = $2
			ORDER BY created_at DESC`, farmID, cycleID)
	} else {
		rows, err = p.db.Query(ctx, `
			SELECT id, title, description, operation_type,
			       production_unit_id, growing_cycle_id,
			       assigned_to, status, priority,
			       due_date, completed_at, created_at
			FROM operations_tasks
			WHERE farm_id = $1
			ORDER BY created_at DESC`, farmID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*taskquery.TaskDTO
	for rows.Next() {
		dto, err := scanTaskDTO(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, dto)
	}
	return result, nil
}

func scanTaskDTO(s interface{ Scan(...any) error }) (*taskquery.TaskDTO, error) {
	var dto taskquery.TaskDTO
	return &dto, s.Scan(
		&dto.ID, &dto.Title, &dto.Description, &dto.OperationType,
		&dto.ProductionUnitID, &dto.GrowingCycleID,
		&dto.AssignedTo, &dto.Status, &dto.Priority,
		&dto.DueDate, &dto.CompletedAt, &dto.CreatedAt,
	)
}
