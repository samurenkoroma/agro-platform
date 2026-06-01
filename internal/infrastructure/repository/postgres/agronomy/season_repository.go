package agronomy

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/season"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type seasonRepo struct {
	tx uow.DB
}

func NewSeasonRepository(tx uow.DB) repository.SeasonRepository {
	return &seasonRepo{
		tx: tx,
	}
}

// Save сохраняет или обновляет сезон
func (r *seasonRepo) Save(ctx context.Context, s *season.Season) error {
	query := `INSERT INTO seasons (id, name, start_date, end_date, status, created_by, owner_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            start_date = EXCLUDED.start_date,
            end_date = EXCLUDED.end_date,
            status = EXCLUDED.status,
            updated_at = EXCLUDED.updated_at
    `

	_, err := r.tx.Exec(ctx, query,
		s.ID,
		s.Name,
		s.StartDate,
		s.EndDate,
		s.Status,
		s.CreatedBy,
		s.OwnerID,
		s.CreatedAt,
		s.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to save season: %w", err)
	}

	return nil
}

func (r *seasonRepo) FindByID(ctx context.Context, id vo.ID) (*season.Season, error) {
	query := `SELECT id, name, start_date, end_date, status, created_by, owner_id, created_at, updated_at
        FROM seasons WHERE id = $1
    `

	var (
		sid       string
		name      string
		startDate time.Time
		endDate   time.Time
		status    string
		createdBy string
		ownerID   string
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.tx.QueryRow(ctx, query, string(id)).Scan(
		&sid, &name, &startDate, &endDate, &ownerID, &status,
		&createdBy, &createdAt, &updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, season.ErrSeasonNotFound
		}
		return nil, fmt.Errorf("failed to find season: %w", err)
	}

	s := &season.Season{
		ID:        vo.ID(sid),
		Name:      name,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    season.SeasonStatus(status),
		CreatedBy: vo.ID(createdBy),
		OwnerID:   vo.ID(ownerID),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return s, nil
}
