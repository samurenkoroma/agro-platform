package getproductionunit

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const QueryName = "spatial.get_production_unit"

type Query struct {
	ID vo.ID
}

func (
	Query,
) QueryName() string {
	return QueryName
}
