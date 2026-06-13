package inventory

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/movement"
	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type movementRepository struct{ db uow.DB }

func NewMovementRepository(db uow.DB) repository.MovementRepository {
	return &movementRepository{db: db}
}

func (r *movementRepository) Save(ctx context.Context, m *movement.Movement) error {
	var refType, refID *string
	if m.Reference != nil {
		rt := string(m.Reference.Type)
		refType = &rt
		refID = &m.Reference.ID
	}
	_, err := r.db.Exec(ctx, `
		INSERT INTO inventory_movements (
			id, farm_id, item_id, type, quantity,
			from_warehouse_id, to_warehouse_id,
			reference_type, reference_id,
			note, timestamp
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		ON CONFLICT (id) DO NOTHING
	`,
		m.ID, m.FarmID, m.ItemID, m.Type, m.Quantity,
		m.FromWarehouseID, m.ToWarehouseID,
		refType, refID,
		m.Note, m.Timestamp,
	)
	return err
}

func (r *movementRepository) GetByID(ctx context.Context, id vo.ID) (*movement.Movement, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, farm_id, item_id, type, quantity,
		       from_warehouse_id, to_warehouse_id,
		       reference_type, reference_id,
		       note, timestamp
		FROM inventory_movements WHERE id = $1`, id)
	return scanMovement(row)
}

func (r *movementRepository) List(ctx context.Context, filter repository.MovementFilter) ([]*movement.Movement, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, farm_id, item_id, type, quantity,
		       from_warehouse_id, to_warehouse_id,
		       reference_type, reference_id,
		       note, timestamp
		FROM inventory_movements
		WHERE farm_id = $1
		ORDER BY timestamp DESC`, filter.FarmID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*movement.Movement
	for rows.Next() {
		m, err := scanMovement(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, m)
	}
	return result, nil
}

func scanMovement(s interface{ Scan(...any) error }) (*movement.Movement, error) {
	var m movement.Movement
	var refType, refID *string
	err := s.Scan(
		&m.ID, &m.FarmID, &m.ItemID, &m.Type, &m.Quantity,
		&m.FromWarehouseID, &m.ToWarehouseID,
		&refType, &refID,
		&m.Note, &m.Timestamp,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, movement.ErrMovementNotFound
		}
		return nil, err
	}
	if refType != nil && refID != nil {
		m.Reference = &movement.Reference{Type: movement.ReferenceType(*refType), ID: *refID}
	}
	return &m, nil
}

var _ repository.MovementRepository = (*movementRepository)(nil)
