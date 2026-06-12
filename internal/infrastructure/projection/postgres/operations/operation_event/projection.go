package operationevent

import (
	"context"
	"encoding/json"

	opquery "github.com/samurenkoroma/agro-platform/internal/application/queries/operations/operation_event"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct{ db uow.DB }

func New(db uow.DB) opquery.Projection { return &projection{db: db} }

func (p *projection) List(ctx context.Context, farmID vo.ID, cycleID *vo.ID) ([]*opquery.OperationDTO, error) {
	var (
		rows interface {
			Next() bool
			Scan(...any) error
			Close()
		}
		err error
	)
	if cycleID != nil {
		rows, err = p.db.Query(ctx, `
			SELECT id, type, production_unit_id, growing_cycle_id, performed_by, payload, timestamp
			FROM operations_events
			WHERE farm_id = $1 AND growing_cycle_id = $2
			ORDER BY timestamp DESC`, farmID, cycleID)
	} else {
		rows, err = p.db.Query(ctx, `
			SELECT id, type, production_unit_id, growing_cycle_id, performed_by, payload, timestamp
			FROM operations_events WHERE farm_id = $1 ORDER BY timestamp DESC`, farmID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*opquery.OperationDTO
	for rows.Next() {
		var dto opquery.OperationDTO
		var payloadJSON []byte
		if err := rows.Scan(&dto.ID, &dto.Type, &dto.ProductionUnitID, &dto.GrowingCycleID,
			&dto.PerformedBy, &payloadJSON, &dto.Timestamp); err != nil {
			return nil, err
		}
		if len(payloadJSON) > 0 {
			json.Unmarshal(payloadJSON, &dto.Payload)
		}
		result = append(result, &dto)
	}
	return result, nil
}
