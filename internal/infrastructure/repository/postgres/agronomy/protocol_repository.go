package agronomy

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	cropprotocol "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop_protocol"
	agronomy "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type protocolRepository struct {
	db uow.DB
}

func NewProtocolRepository(db uow.DB) agronomy.CropProtocolRepository {
	return &protocolRepository{db: db}
}

func (r *protocolRepository) GetByID(ctx context.Context, id vo.ID) (*cropprotocol.CropProtocol, error) {
	query := `
SELECT
	id,
	crop_id,
	name,
	growing_method,
	description,
	metadata,
	created_at,
	updated_at,
	archived_at
FROM crop_protocols
WHERE id=$1
`

	root := &cropprotocol.CropProtocol{}

	err := r.db.QueryRow(ctx, query, id).Scan(scanCropProtocol(root)...)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	stages, err := r.loadStages(ctx, root.ID)
	if err != nil {
		return nil, err
	}

	root.StageProfiles = stages

	return root, nil
}

func (r *protocolRepository) GetByCrop(ctx context.Context, cropID vo.ID) ([]*cropprotocol.CropProtocol, error) {
	query := `
SELECT
	id,
	crop_id,
	name,
	growing_method,
	description,
	metadata,
	created_at,
	updated_at,
	archived_at
FROM crop_protocols
WHERE crop_id=$1
ORDER BY name
`

	rows, err := r.db.Query(ctx, query, cropID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*cropprotocol.CropProtocol, 0)

	for rows.Next() {
		root := &cropprotocol.CropProtocol{}

		if err := rows.Scan(scanCropProtocol(root)...); err != nil {
			return nil, err
		}

		stages, err := r.loadStages(ctx, root.ID)
		if err != nil {
			return nil, err
		}

		root.StageProfiles = stages

		result = append(result, root)
	}

	return result, nil
}
func (r *protocolRepository) Save(ctx context.Context, root *cropprotocol.CropProtocol) error {
	query := `INSERT INTO 
    crop_protocols(id,crop_id,name,growing_method,description,metadata,created_at,updated_at,archived_at)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
				ON CONFLICT(id)
				DO UPDATE SET
					name=excluded.name,
					growing_method=excluded.growing_method,
					description=excluded.description,
					metadata=excluded.metadata,
					updated_at=excluded.updated_at,
					archived_at=excluded.archived_at`

	_, err := r.db.Exec(
		ctx,
		query,
		root.ID,
		root.CropID,
		root.Name,
		root.Description,
		root.Metadata,
		root.CreatedAt,
		root.UpdatedAt,
		root.ArchivedAt,
	)

	if err != nil {
		return err
	}

	for _, stage := range root.StageProfiles {

		stageQuery := `INSERT INTO crop_protocol_stage_profiles(
	id,protocol_id,crop_stage_id,climate,lighting,irrigation,nutrition,water,vpd,order_index,metadata
	)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`

		_, err = r.db.Exec(
			ctx,
			stageQuery,
			stage.ID,
			root.ID,
			stage.CropStageID,
			stage.Climate,
			stage.Lighting,
			stage.Irrigation,
			stage.Nutrition,
			stage.Water,
			stage.VPD,
			0,
			root.Metadata,
		)

		if err != nil {
			return err
		}
	}

	return err
}

func (r *protocolRepository) Exists(ctx context.Context, id vo.ID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1	FROM crop_protocols WHERE id=$1)`
	var exists bool

	err := r.db.QueryRow(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *protocolRepository) loadStages(ctx context.Context, protocolID vo.ID) ([]cropprotocol.StageProfile, error) {
	query := `
SELECT
	id,
	crop_stage_id,
	climate,
	lighting,
	irrigation,
	nutrition,
	water,
	vpd
FROM crop_protocol_stage_profiles
WHERE protocol_id=$1
ORDER BY order_index
`

	rows, err := r.db.Query(ctx, query, protocolID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]cropprotocol.StageProfile, 0)

	for rows.Next() {
		var stage cropprotocol.StageProfile

		err = rows.Scan(scanStage(&stage)...)
		if err != nil {
			return nil, err
		}

		result = append(result, stage)
	}

	return result, nil
}
