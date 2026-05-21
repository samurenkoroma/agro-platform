package crop

import (
	"context"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

func (
	r *Repository,
) Exists(
	ctx context.Context,

	id vo.ID,
) (
	bool,
	error,
) {

	query := `
SELECT EXISTS(
SELECT 1
FROM crops
WHERE id=$1
)
`

	var exists bool

	err :=
		r.db.QueryRow(
			ctx,
			query,
			id,
		).
			Scan(
				&exists,
			)

	if err != nil {
		return false,
			err
	}

	return exists,
		nil
}
