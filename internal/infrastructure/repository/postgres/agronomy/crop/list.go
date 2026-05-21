package crop

import (
	"context"

	entity "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
)

func (
	r *Repository,
) GetAll(
	ctx context.Context,
) (
	[]*entity.Crop,
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

ORDER BY name
`

	rows,
		err :=
		r.db.Query(
			ctx,
			query,
		)

	if err != nil {
		return nil,
			err
	}

	defer rows.Close()

	result :=
		make(
			[]*entity.Crop,
			0,
		)

	for rows.Next() {

		root :=
			&entity.Crop{}

		err =
			rows.Scan(
				scanTarget(
					root,
				)...,
			)

		if err != nil {
			return nil,
				err
		}

		result =
			append(
				result,
				root,
			)
	}

	return result,
		nil
}
