package crop

import (
	"context"

	"github.com/jackc/pgx/v5"

	entity "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/crop"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func (
	r *Repository,
) GetByID(
	ctx context.Context,

	id vo.ID,
) (
	*entity.Crop,
	error,
) {

	query := `
SELECT

id,
name,
scientific_name,
category,
metadata,
created_at,
updated_at

FROM crops

WHERE id=$1
`

	row :=
		r.db.QueryRow(
			ctx,
			query,
			id,
		)

	root :=
		&entity.Crop{}

	err :=
		row.Scan(
			scanTarget(
				root,
			)...,
		)

	if err != nil {

		if err ==
			pgx.ErrNoRows {

			return nil,
				nil
		}

		return nil,
			err
	}

	return root,
		nil
}
