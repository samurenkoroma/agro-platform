package agronomy

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	stage "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop_stage"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Repository struct {
	db uow.DB
}

func NewCropStageRepository(db uow.DB) repository.CropStageRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Save(ctx context.Context, root *stage.CropStage) error {
	query := `INSERT INTO 
    crop_stages(id,crop_id,code,name,bbch,order_index,duration_days,metadata,created_at,updated_at,archived_at)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		ON CONFLICT(id)
		DO UPDATE SET
	code=excluded.code,
	name=excluded.name,
	bbch=excluded.bbch,
	order_index=excluded.order_index,
	duration_days=excluded.duration_days,
	metadata=excluded.metadata,
	updated_at=excluded.updated_at,
	archived_at=excluded.archived_at`

	_, err := r.db.Exec(
		ctx,
		query,
		root.ID,
		root.CropID,
		root.Code,
		root.Name,
		root.BBCH,
		root.Order,
		root.DurationDays,
		root.Metadata,
		root.CreatedAt,
		root.UpdatedAt,
		root.ArchivedAt,
	)

	return err
}

func (r *Repository) GetByID(ctx context.Context, id vo.ID) (*stage.CropStage, error) {
	query := `SELECT id,crop_id,code,name,bbch,order_index,duration_days,metadata,created_at,updated_at,archived_at
				FROM crop_stages
				WHERE id=$1`

	root := &stage.CropStage{}

	err := r.db.QueryRow(ctx, query, id).Scan(scanCropStage(root)...)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return root, nil
}

func (r *Repository) GetByCrop(ctx context.Context, cropID vo.ID) ([]*stage.CropStage, error) {
	query := `SELECT id,crop_id,code,name,bbch,order_index,duration_days,metadata,created_at,updated_at,archived_at
				FROM crop_stages
				WHERE crop_id=$1
				ORDER BY order_index`

	rows, err := r.db.Query(ctx, query, cropID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*stage.CropStage, 0)

	for rows.Next() {
		root := &stage.CropStage{}

		if err := rows.Scan(scanCropStage(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *Repository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM crop_stages WHERE id=$1)`

	var exists bool

	err := r.db.QueryRow(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
