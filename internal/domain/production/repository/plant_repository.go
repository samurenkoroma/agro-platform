package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	plant "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/plant"
)

type PlantRepository interface {
	Save(aggregate *plant.Aggregate) error
	GetByID(id vo.ID) (*plant.Aggregate, error)
	GetByCycle(cycleID vo.ID) ([]*plant.Aggregate, error)
	Delete(id vo.ID) error
}
