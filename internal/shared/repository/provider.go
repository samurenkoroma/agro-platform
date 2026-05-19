package repository

import (
	environmentrepo "github.com/samurenkoroma/agro-platform/internal/domain/environment/repository"
	operationsrepo "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	productionrepo "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"
	spatialrepo "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
)

type Provider interface {
	ProductionUnit() spatialrepo.ProductionUnitRepository

	GrowingCycle() productionrepo.GrowingCycleRepository

	Plant() productionrepo.PlantRepository

	Task() operationsrepo.TaskRepository

	Sensor() environmentrepo.SensorRepository
}
