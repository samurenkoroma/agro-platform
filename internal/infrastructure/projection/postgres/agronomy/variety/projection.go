package variety

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/variety"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
)

type projection struct {
	db uow.DB
}

func New(db uow.DB) variety.Projection {
	return &projection{
		db: db,
	}
}

func (p projection) Get(ctx context.Context, id string) (*variety.Detail, error) {
	query := `SELECT  id,name FROM varieties WHERE id=$1`

	var result variety.Detail

	err := p.db.QueryRow(ctx, query, id).Scan(
		&result.ID,
		&result.Name,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (p projection) List(ctx context.Context, filter variety.ListFilter) ([]variety.ListItem, error) {

	query := `SELECT id,name FROM varieties ORDER BY name= $1 `
	rows, err := p.db.Query(ctx, query, filter.CropKey)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]variety.ListItem, 0)

	for rows.Next() {

		var item variety.ListItem

		err = rows.Scan(&item.ID, &item.Name, &item.SpeciesName)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}
