package getproductionunittree

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const QueryName = "spatial.get_production_unit_tree"

type Query struct {
	RootID vo.ID
}

func (
	Query,
) QueryName() string {
	return QueryName
}
