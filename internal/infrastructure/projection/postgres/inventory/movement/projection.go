package movement

import (
	"context"

	movquery "github.com/samurenkoroma/agro-platform/internal/application/queries/inventory/movement"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct{ db uow.DB }

func New(db uow.DB) movquery.Projection { return &projection{db: db} }

func (p *projection) List(ctx context.Context, farmID vo.ID, itemID *vo.ID) ([]*movquery.MovementDTO, error) {
	var (
		rows interface {
			Next() bool
			Scan(...any) error
			Close()
		}
		err error
	)
	if itemID != nil {
		rows, err = p.db.Query(ctx, `
			SELECT id, item_id, type, quantity,
			       from_warehouse_id, to_warehouse_id,
			       reference_type, reference_id, note, timestamp
			FROM inventory_movements
			WHERE farm_id = $1 AND item_id = $2
			ORDER BY timestamp DESC`, farmID, itemID)
	} else {
		rows, err = p.db.Query(ctx, `
			SELECT id, item_id, type, quantity,
			       from_warehouse_id, to_warehouse_id,
			       reference_type, reference_id, note, timestamp
			FROM inventory_movements
			WHERE farm_id = $1
			ORDER BY timestamp DESC`, farmID)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*movquery.MovementDTO
	for rows.Next() {
		var dto movquery.MovementDTO
		if err := rows.Scan(
			&dto.ID, &dto.ItemID, &dto.Type, &dto.Quantity,
			&dto.FromWarehouseID, &dto.ToWarehouseID,
			&dto.ReferenceType, &dto.ReferenceID, &dto.Note, &dto.Timestamp,
		); err != nil {
			return nil, err
		}
		result = append(result, &dto)
	}
	return result, nil
}
