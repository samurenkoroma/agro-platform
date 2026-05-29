package agronomy

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	entity "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/variety"
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type varietyRepository struct {
	db uow.DB
}

func NewVarietyRepository(db uow.DB) repository.VarietyRepository {
	return &varietyRepository{db: db}
}

func (r *varietyRepository) Exists(ctx context.Context, name string, cropId vo.ID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM varieties WHERE name=$1 AND crop_id=$2)`

	var exists bool

	err := r.db.QueryRow(ctx, query, name, cropId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *varietyRepository) GetByID(ctx context.Context, id vo.ID) (*entity.Variety, error) {
	query := `SELECT id,crop_id,name,breeder,maturity,growth,spacing,tolerance,metadata,created_at,updated_at
				FROM varieties
				WHERE id=$1`

	root := &entity.Variety{}
	err := r.db.QueryRow(ctx, query, id).Scan(scanVariety(root)...)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return root, nil
}

func (r *varietyRepository) GetByCrop(ctx context.Context, cropID vo.ID) ([]*entity.Variety, error) {
	query := `SELECT id,crop_id,name,breeder, maturity,growth,spacing, harvest, yield_profile, tolerance,metadata,created_at,updated_at, archived_at
				FROM varieties
				WHERE crop_id=$1
				ORDER BY name`

	rows, err := r.db.Query(ctx, query, cropID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.Variety, 0)

	for rows.Next() {
		root := &entity.Variety{}
		if err := rows.Scan(scanVariety(root)...); err != nil {
			return nil, err
		}

		result = append(result, root)
	}

	return result, nil
}

func (r *varietyRepository) Save(ctx context.Context, root *entity.Variety) error {
	query := `INSERT INTO varieties(
                      id,crop_id,name,breeder,
                      maturity,growth,spacing,harvest,yield_profile,tolerance,
                      characteristics,base_temperature,max_temperature,water_requirement,light_requirement,phenophase_gdd,
                      metadata,created_at,updated_at
                      )
				VALUES(
				       $1,$2,$3,$4,
				       $5,$6,$7,$8,$9,$10,
				       $11,$12,$13, $14, $15, $16,
				       $17, $18, $19)
				ON CONFLICT(id)
				DO UPDATE SET
					name=excluded.name,
					breeder=excluded.breeder,
					maturity=excluded.maturity,
					growth=excluded.growth,
					spacing=excluded.spacing,
					tolerance=excluded.tolerance,
					metadata=excluded.metadata,
					harvest=excluded.harvest,
					yield_profile=excluded.yield_profile,
					updated_at=excluded.updated_at`

	_, err := r.db.Exec(ctx, query,
		root.ID, root.CropID, root.Name, root.Breeder,

		root.Maturity, root.Growth, root.Spacing, root.Harvest, root.Yield, root.Tolerance,
		root.Characteristics, root.BaseTemperature, root.MaxTemperature, root.WaterRequirement, root.LightRequirement, root.PhenophaseGDD,
		root.Metadata, root.CreatedAt, root.UpdatedAt,
	)

	return err
}
