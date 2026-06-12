package timeline

import (
	"context"
	"encoding/json"

	tlquery "github.com/samurenkoroma/agro-platform/internal/application/queries/operations/timeline"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct{ db uow.DB }

func New(db uow.DB) tlquery.Projection { return &projection{db: db} }

func (p *projection) Get(ctx context.Context, farmID vo.ID, cycleID *vo.ID) (*tlquery.TimelineDTO, error) {
	var row interface{ Scan(...any) error }
	if cycleID != nil {
		row = p.db.QueryRow(ctx, `
			SELECT id, items FROM operations_timelines
			WHERE farm_id = $1 AND growing_cycle_id = $2 LIMIT 1`, farmID, cycleID)
	} else {
		row = p.db.QueryRow(ctx, `
			SELECT id, items FROM operations_timelines
			WHERE farm_id = $1 AND growing_cycle_id IS NULL LIMIT 1`, farmID)
	}
	var dto tlquery.TimelineDTO
	var itemsJSON []byte
	if err := row.Scan(&dto.ID, &itemsJSON); err != nil {
		return &tlquery.TimelineDTO{Items: []tlquery.TimelineItemDTO{}}, nil
	}
	if len(itemsJSON) > 0 {
		json.Unmarshal(itemsJSON, &dto.Items)
	}
	return &dto, nil
}
