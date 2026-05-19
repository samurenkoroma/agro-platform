package agronomy

import (
	"github.com/samurenkoroma/agro-platform/internal/domain/agronomy/aggregate/variety"
	agronomyrepo "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type varietyRepository struct {
	items map[vo.ID]*variety.Variety
}

func NewVarietyRepository() agronomyrepo.VarietyRepository {
	return &varietyRepository{items: make(map[vo.ID]*variety.Variety)}
}
