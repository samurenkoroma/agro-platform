package season

import (
	"context"
	"fmt"
	"time"

	"github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/season"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
)

type catalogProjection struct {
	db uow.DB
}

func (c catalogProjection) List(ctx context.Context, filter season.SeasonFilter) ([]season.SeasonItem, error) {
	query := `SELECT id, name, start_date, end_date, status FROM seasons
WHERE owner_id = $1 ORDER BY start_date DESC`

	rows, err := c.db.Query(ctx, query, filter.OwnerId)
	if err != nil {
		return nil, fmt.Errorf("failed to query seasons: %w", err)
	}
	defer rows.Close()
	items := make([]season.SeasonItem, 0)
	for rows.Next() {
		var item season.SeasonItem = season.SeasonItem{
			Id:        "",
			Name:      "",
			StartDate: time.Time{},
			EndDate:   time.Time{},
			Status:    "",
			Statistics: season.Statistics{
				TotalPlans:     0,
				ActivePlans:    0,
				TotalArea:      0,
				CompletedPlans: 0,
				TotalHarvest:   0,
				Crops:          []season.StatItem{},
			},
			PlantingArea: []season.SeasonUnits{},
			Weather:      season.Weather{},
		}
		if err := rows.Scan(&item.Id, &item.Name, &item.StartDate, &item.EndDate, &item.Status); err != nil {
			return nil, fmt.Errorf("failed to scan season: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate seasons: %w", err)
	}
	return items, nil
}

func New(db uow.DB) season.Projection {
	return &catalogProjection{db: db}
}
