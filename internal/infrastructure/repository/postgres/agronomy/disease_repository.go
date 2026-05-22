package agronomy

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/disease"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type diseaseRepository struct {
	db uow.DB
}

func New(db uow.DB) repository.DiseaseRepository {
	return &diseaseRepository{
		db: db,
	}
}

func (r *diseaseRepository) Save(ctx context.Context, root *disease.Disease) error {
	query := `
INSERT INTO diseases(
	id,
	name,
	scientific_name,
	pathogen_type,
	hosts,
	symptoms,
	description,
	metadata,
	created_at,
	updated_at,
	archived_at
)
VALUES(
	$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11
)
ON CONFLICT(id)
DO UPDATE SET
	name=excluded.name,
	scientific_name=excluded.scientific_name,
	pathogen_type=excluded.pathogen_type,
	hosts=excluded.hosts,
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
		root.ScientificName,

		root.PathogenType,

		root.Hosts,
		root.Symptoms,

		root.Description,

		root.Metadata,

		root.CreatedAt,
		root.UpdatedAt,
		root.ArchivedAt,
	)

	return err
}

func (r *diseaseRepository) GetByID(ctx context.Context, id vo.ID) (*disease.Disease, error) {
	query := `
SELECT
	id,
	name,
	scientific_name,
	pathogen_type,
	hosts,
	symptoms,
	description,
	metadata,
	created_at,
	updated_at,
	archived_at
FROM diseases
WHERE id=$1
`

	root := &disease.Disease{}

	err := r.db.QueryRow(ctx, query, id).Scan(scanDisease(root)...)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return root, nil
}

func (r *diseaseRepository) GetAll(ctx context.Context) ([]*disease.Disease, error) {
	query := `
SELECT
	id,
	name,
	scientific_name,
	pathogen_type,
	hosts,
	symptoms,
	description,
	metadata,
	created_at,
	updated_at,
	archived_at
FROM diseases
ORDER BY name
`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*disease.Disease, 0)

	for rows.Next() {
		root := &disease.Disease{}

		if err := rows.Scan(scanDisease(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *diseaseRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	query := `
SELECT EXISTS(
	SELECT 1
	FROM diseases
	WHERE id=$1
)
`

	var exists bool

	err := r.db.QueryRow(ctx, query, id).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func NewDiseaseRepository(db uow.DB) repository.DiseaseRepository {
	return &diseaseRepository{db: db}
}
