package repository

import (
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

	plant "github.com/samurenkoroma/agro-platform/internal/domain/production/aggregate/plant"
)

type PlantRepository interface {
	Save(aggregate *plant.Plant) error
	GetByID(id vo.ID) (*plant.Plant, error)
	GetByCycle(cycleID vo.ID) ([]*plant.Plant, error)
	Delete(id vo.ID) error
}
