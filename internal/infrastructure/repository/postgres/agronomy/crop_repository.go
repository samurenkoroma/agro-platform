package agronomy

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	entity "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type cropRepository struct {
	db uow.DB
}

func NewCropRepository(db uow.DB) repository.CropRepository {
	return &cropRepository{db: db}
}

func (r *cropRepository) GetByID(ctx context.Context, id vo.ID) (*entity.Crop, error) {
	query := `SELECT id,name,scientific_name,category,metadata,created_at,updated_at
				FROM crops
				WHERE id=$1`

	row := r.db.QueryRow(ctx, query, id)
	root := &entity.Crop{}
	err := row.Scan(scanCrop(root)...)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return root, nil
}

func (r *cropRepository) GetAll(ctx context.Context) ([]*entity.Crop, error) {
	query := `SELECT id,name,scientific_name,category,metadata,created_at,updated_at
				FROM crops
				ORDER BY name`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.Crop, 0)

	for rows.Next() {
		root := &entity.Crop{}

		err = rows.Scan(scanCrop(root)...)

		if err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *cropRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM crops WHERE id=$1)`

	var exists bool

	err := r.db.QueryRow(ctx, query, id).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *cropRepository) Save(ctx context.Context, root *entity.Crop) error {
	query := `INSERT INTO crops(id,name,scientific_name,category,metadata,created_at,updated_at)
				VALUES($1,$2,$3,$4,$5,$6,$7)
				ON CONFLICT(id)
				DO UPDATE SET
			name=excluded.name,
			scientific_name=excluded.scientific_name,
			category=excluded.category,
			metadata=excluded.metadata,
			updated_at=excluded.updated_at`

	_, err := r.db.Exec(ctx, query,
		root.ID, root.Name, root.ScientificName, root.Category, root.Metadata, root.CreatedAt, root.UpdatedAt)

	return err
}
