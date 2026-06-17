package crop

import (
	"context"

	"github.com/samurenkoroma/agro-platform/internal/application/queries/agronomy/crop"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
)

type projection struct {
	db uow.DB
}

func New(db uow.DB) crop.Projection {
	return &projection{
		db: db,
	}
}

func (p projection) Get(ctx context.Context, id string) (*crop.Detail, error) {

	query := `SELECT id,name,category, family,scientific_name, imageurl FROM crops WHERE id=$1`

	var result crop.Detail

	err := p.db.QueryRow(ctx, query, id).Scan(
		&result.ID,
		&result.Name,
		&result.Category,
		&result.Family,
		&result.ScientificName,
		&result.ImageUrl,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (p projection) List(ctx context.Context, filter crop.ListFilter) ([]crop.ListItem, error) {

	query := `SELECT id,name,category,family, scientific_name, imageurl 
FROM crops 
WHERE (
    COALESCE(array_length($1::text[], 1), 0) = 0
    OR category = ANY($1::text[])
)
ORDER BY name`
	rows, err := p.db.Query(ctx, query, filter.Category)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]crop.ListItem, 0)

	for rows.Next() {

		var item crop.ListItem

		err = rows.Scan(&item.ID, &item.Name, &item.Category, &item.Family, &item.ScientificName, &item.ImageUrl)

		if err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}
