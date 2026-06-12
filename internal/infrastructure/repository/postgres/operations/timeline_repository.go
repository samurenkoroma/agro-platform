package operations

import (
	"context"
	"encoding/json"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/aggregate/timeline"
	"github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type timelineRepository struct{ db uow.DB }

func NewTimelineRepository(db uow.DB) repository.TimeLineRepository {
	return &timelineRepository{db: db}
}

func (r *timelineRepository) Save(ctx context.Context, t *timeline.Timeline) error {
	itemsJSON, err := json.Marshal(t.Items)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, `
		INSERT INTO operations_timelines (
			id, farm_id, production_unit_id, growing_cycle_id, items, created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (id) DO UPDATE SET
			items      = EXCLUDED.items,
			updated_at = EXCLUDED.updated_at
	`, t.ID, t.FarmID, t.ProductionUnitID, t.GrowingCycleID, itemsJSON, t.CreatedAt, t.UpdatedAt)
	return err
}

func (r *timelineRepository) GetByID(ctx context.Context, id vo.ID) (*timeline.Timeline, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, farm_id, production_unit_id, growing_cycle_id, items, created_at, updated_at
		FROM operations_timelines WHERE id = $1`, id)
	return scanTimeline(row)
}

func (r *timelineRepository) GetByOwner(ctx context.Context, farmID vo.ID, cycleID *vo.ID) (*timeline.Timeline, error) {
	var row interface{ Scan(...any) error }
	if cycleID != nil {
		row = r.db.QueryRow(ctx, `
			SELECT id, farm_id, production_unit_id, growing_cycle_id, items, created_at, updated_at
			FROM operations_timelines
			WHERE farm_id = $1 AND growing_cycle_id = $2
			LIMIT 1`, farmID, cycleID)
	} else {
		row = r.db.QueryRow(ctx, `
			SELECT id, farm_id, production_unit_id, growing_cycle_id, items, created_at, updated_at
			FROM operations_timelines
			WHERE farm_id = $1 AND growing_cycle_id IS NULL
			LIMIT 1`, farmID)
	}
	t, err := scanTimeline(row)
	if err != nil {
		return nil, nil
	}
	return t, nil
}

func scanTimeline(s interface{ Scan(...any) error }) (*timeline.Timeline, error) {
	var t timeline.Timeline
	var itemsJSON []byte
	if err := s.Scan(&t.ID, &t.FarmID, &t.ProductionUnitID, &t.GrowingCycleID, &itemsJSON, &t.CreatedAt, &t.UpdatedAt); err != nil {
		return nil, err
	}
	if len(itemsJSON) > 0 {
		if err := json.Unmarshal(itemsJSON, &t.Items); err != nil {
			return nil, err
		}
	}
	return &t, nil
}

var _ repository.TimeLineRepository = (*timelineRepository)(nil)
