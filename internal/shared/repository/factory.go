package repository

import (
	//analyticsrepo "github.com/samurenkoroma/agro-platform/internal/domain/analytics/repository"

	agronomyrepo "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"

	//automationrepo "github.com/samurenkoroma/agro-platform/internal/domain/automation/repository"

	environmentrepo "github.com/samurenkoroma/agro-platform/internal/domain/environment/repository"

	//inventoryrepo "github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"

	operationsrepo "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"

	productionrepo "github.com/samurenkoroma/agro-platform/internal/domain/production/repository"

	spatialrepo "github.com/samurenkoroma/agro-platform/internal/domain/spatial/repository"
)

type AnalyticsProvider interface {
	//analyticsrepo.
}
type AutomationProvider interface {
	//automationrepo.
}
type InventoryProvider interface {
	//inventoryrepo.
}
type EnvironmentProvider interface {
	Sensors() environmentrepo.SensorRepository
}
type OperationsProvider interface {
	Tasks() operationsrepo.TaskRepository
}
type SpatialProvider interface {
	ProductionUnits() spatialrepo.ProductionUnitRepository
}
type ProductionProvider interface {
	GrowingCycles() productionrepo.GrowingCycleRepository

	Plants() productionrepo.PlantRepository

	Slots() productionrepo.SlotRepository

	Substrates() productionrepo.SubstrateRepository

	Harvests() productionrepo.HarvestRepository

	Yields() productionrepo.YieldRepository
}

type AgronomyProvider interface {
	Crops() agronomyrepo.CropRepository

	Varieties() agronomyrepo.VarietyRepository

	Protocols() agronomyrepo.ProtocolRepository

	Diseases() agronomyrepo.DiseaseRepository

	StressProfiles() agronomyrepo.StressRepository
}
