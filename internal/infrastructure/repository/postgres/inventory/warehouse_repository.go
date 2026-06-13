package inventory

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/aggregate/warehouse"
	"github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type warehouseRepository struct{ db uow.DB }

func NewWarehouseRepository(db uow.DB) repository.WarehouseRepository {
	return &warehouseRepository{db: db}
}

func (r *warehouseRepository) Save(ctx context.Context, w *warehouse.Warehouse) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO inventory_warehouses (
			id, farm_id, name, code, created_at, updated_at, archived_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (id) DO UPDATE SET
			name        = EXCLUDED.name,
			code        = EXCLUDED.code,
			updated_at  = EXCLUDED.updated_at,
			archived_at = EXCLUDED.archived_at
	`, w.ID, w.FarmID, w.Name, w.Code, w.CreatedAt, w.UpdatedAt, w.ArchivedAt)
	return err
}

func (r *warehouseRepository) GetByID(ctx context.Context, id vo.ID) (*warehouse.Warehouse, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, farm_id, name, code, created_at, updated_at, archived_at
		FROM inventory_warehouses WHERE id = $1`, id)
	return scanWarehouse(row)
}

func (r *warehouseRepository) List(ctx context.Context, farmID vo.ID) ([]*warehouse.Warehouse, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, farm_id, name, code, created_at, updated_at, archived_at
		FROM inventory_warehouses
		WHERE farm_id = $1 AND archived_at IS NULL
		ORDER BY name ASC`, farmID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*warehouse.Warehouse
	for rows.Next() {
		w, err := scanWarehouse(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, w)
	}
	return result, nil
}

func scanWarehouse(s interface{ Scan(...any) error }) (*warehouse.Warehouse, error) {
	var w warehouse.Warehouse
	err := s.Scan(&w.ID, &w.FarmID, &w.Name, &w.Code, &w.CreatedAt, &w.UpdatedAt, &w.ArchivedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, warehouse.ErrWarehouseNotFound
		}
		return nil, err
	}
	return &w, nil
}

var _ repository.WarehouseRepository = (*warehouseRepository)(nil)
