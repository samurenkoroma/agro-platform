package crop

import (
	"context"

	entity "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/crop"
)

func (
	r *Repository,
) Save(
	ctx context.Context,

	root *entity.Crop,
) error {

	query := `
INSERT INTO crops(
id,
name,
scientific_name,
category,
metadata,
created_at,
updated_at
)

VALUES(
$1,$2,$3,$4,$5,$6,$7
)

ON CONFLICT(id)

DO UPDATE SET

name=excluded.name,

scientific_name=
excluded.scientific_name,

category=
excluded.category,

metadata=
excluded.metadata,

updated_at=
excluded.updated_at
`

	_,
		err :=
		r.db.Exec(
			ctx,
			query,

			root.ID,

			root.Name,

			root.ScientificName,

			root.Category,

			root.Metadata,

			root.CreatedAt,

			root.UpdatedAt,
		)

	return err
}
