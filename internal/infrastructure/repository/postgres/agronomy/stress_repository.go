package agronomy

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/stress"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type stressRepository struct {
	db uow.DB
}

func NewStressRepository(db uow.DB) repository.StressRepository {
	return &stressRepository{
		db: db,
	}
}

func (r *stressRepository) Save(
	ctx context.Context,
	root *stress.Stress,
) error {

	query := `
INSERT INTO stresses(
    id,
    name,
    type,
    triggers,
    symptoms,
    description,
    metadata,
    created_at,
    updated_at,
    archived_at
)
VALUES(
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10
)
ON CONFLICT(id)
DO UPDATE SET
    name=excluded.name,
    type=excluded.type,
    triggers=excluded.triggers,
    symptoms=excluded.symptoms,
    description=excluded.description,
    metadata=excluded.metadata,
    updated_at=excluded.updated_at,
    archived_at=excluded.archived_at
`

	_, err := r.db.Exec(
		ctx,
		query,

		root.ID,
		root.Name,
		root.Type,

		root.Triggers,
		root.Symptoms,

		root.Description,

		root.Metadata,

		root.CreatedAt,
		root.UpdatedAt,
		root.ArchivedAt,
	)

	return err
}

func (r *stressRepository) GetByID(
	ctx context.Context,
	id vo.ID,
) (*stress.Stress, error) {

	query := `
SELECT
	id,
	name,
	type,
	description,
	triggers,
	symptoms,
	metadata,
	created_at,
	updated_at,
	archived_at
FROM stresses
WHERE id=$1
`

	root := &stress.Stress{}

	err := r.db.
		QueryRow(ctx, query, id).
		Scan(scanStress(root)...)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return root, nil
}

func (r *stressRepository) GetAll(
	ctx context.Context,
) ([]*stress.Stress, error) {

	query := `
SELECT
	id,
	name,
	type,
	description,
	triggers,
	symptoms,
	metadata,
	created_at,
	updated_at,
	archived_at
FROM stresses
ORDER BY name
`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*stress.Stress, 0)

	for rows.Next() {
		root := &stress.Stress{}

		if err := rows.Scan(
			scanStress(root)...,
		); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *stressRepository) Exists(
	ctx context.Context,
	id vo.ID,
) (bool, error) {

	query := `
SELECT EXISTS(
	SELECT 1
	FROM stresses
	WHERE id=$1
)
`

	var exists bool

	err := r.db.
		QueryRow(ctx, query, id).
		Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
