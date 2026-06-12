package operations

import (
	"context"
	"encoding/json"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	operationevent "github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/operation_event"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type operationRepository struct{ db uow.DB }

func NewOperationRepository(db uow.DB) repository.OperationRepository {
	return &operationRepository{db: db}
}

func (r *operationRepository) Save(ctx context.Context, e *operationevent.OperationEvent) error {
	payloadJSON, err := json.Marshal(e.Payload)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, `
		INSERT INTO operations_events (
			id, farm_id, type, production_unit_id, growing_cycle_id,
			performed_by, payload, timestamp
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (id) DO NOTHING
	`, e.ID, e.FarmID, e.Type, e.ProductionUnitID, e.GrowingCycleID,
		e.PerformedBy, payloadJSON, e.Timestamp)
	return err
}

func (r *operationRepository) GetByID(ctx context.Context, id vo.ID) (*operationevent.OperationEvent, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, farm_id, type, production_unit_id, growing_cycle_id,
		       performed_by, payload, timestamp
		FROM operations_events WHERE id = $1`, id)
	return scanOperation(row)
}

func (r *operationRepository) List(ctx context.Context, filter repository.OperationFilter) ([]*operationevent.OperationEvent, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, farm_id, type, production_unit_id, growing_cycle_id,
		       performed_by, payload, timestamp
		FROM operations_events
		WHERE farm_id = $1
		ORDER BY timestamp DESC`, filter.FarmID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*operationevent.OperationEvent
	for rows.Next() {
		e, err := scanOperation(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}

func scanOperation(s interface{ Scan(...any) error }) (*operationevent.OperationEvent, error) {
	var e operationevent.OperationEvent
	var payloadJSON []byte
	if err := s.Scan(&e.ID, &e.FarmID, &e.Type, &e.ProductionUnitID, &e.GrowingCycleID,
		&e.PerformedBy, &payloadJSON, &e.Timestamp); err != nil {
		return nil, err
	}
	if len(payloadJSON) > 0 {
		if err := json.Unmarshal(payloadJSON, &e.Payload); err != nil {
			return nil, err
		}
	}
	return &e, nil
}

var _ repository.OperationRepository = (*operationRepository)(nil)
