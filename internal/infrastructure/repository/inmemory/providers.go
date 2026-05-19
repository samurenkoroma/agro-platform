package inmemory

import (
	agronomyrepo "github.com/samurenkoroma/agro-platform/internal/domain/agronomy/repository"
	analyticsrepo "github.com/samurenkoroma/agro-platform/internal/domain/analytics/repository"
	automationrepo "github.com/samurenkoroma/agro-platform/internal/domain/automation/repository"
	environmentrepo "github.com/samurenkoroma/agro-platform/internal/domain/environment/repository"
	inventoryrepo "github.com/samurenkoroma/agro-platform/internal/domain/inventory/repository"
	operationsrepo "github.com/samurenkoroma/agro-platform/internal/domain/operations/repository"
	productioninfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/production"
	spatialinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/spatial"

	agronomyinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/agronomy"
	analyticsinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/analytics"
	automationinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/automation"
	environmentinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/environment"
	inventoryinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/inventory"
	operationsinfra "github.com/samurenkoroma/agro-platform/internal/infrastructure/repository/inmemory/operations"
)

//
// PRODUCTION
//

type ProductionProvider struct {
	growingCycles *productioninfra.GrowingCycleRepository
	plants        *productioninfra.PlantRepository
}

func NewProductionProvider() *ProductionProvider {
	return &ProductionProvider{
		growingCycles: productioninfra.NewGrowingCycleRepository(),
		plants:        productioninfra.NewPlantRepository(),
	}
}

func (p *ProductionProvider) GrowingCycles() *productioninfra.GrowingCycleRepository {
	return p.growingCycles
}

func (p *ProductionProvider) Plants() *productioninfra.PlantRepository {
	return p.plants
}

//
// SPATIAL
//

type SpatialProvider struct {
	productionUnits *spatialinfra.ProductionUnitRepository
}

func NewSpatialProvider() *SpatialProvider {
	return &SpatialProvider{
		productionUnits: spatialinfra.NewProductionUnitRepository(),
	}
}

func (p *SpatialProvider) ProductionUnits() *spatialinfra.ProductionUnitRepository {
	return p.productionUnits
}

//
// AGRONOMY
//

type AgronomyProvider struct {
	crops     agronomyrepo.CropRepository
	varieties agronomyrepo.VarietyRepository
	protocols agronomyrepo.ProtocolRepository
	diseases  agronomyrepo.DiseaseRepository
	stress    agronomyrepo.StressRepository
}

func NewAgronomyProvider() *AgronomyProvider {
	return &AgronomyProvider{
		crops:     agronomyinfra.NewCropRepository(),
		varieties: agronomyinfra.NewVarietyRepository(),
		protocols: agronomyinfra.NewProtocolRepository(),
		diseases:  agronomyinfra.NewDiseaseRepository(),
		stress:    agronomyinfra.NewStressRepository(),
	}
}

func (p *AgronomyProvider) Crops() agronomyrepo.CropRepository {
	return p.crops
}
func (p *AgronomyProvider) Varieties() agronomyrepo.VarietyRepository {
	return p.varieties
}
func (p *AgronomyProvider) Protocols() agronomyrepo.ProtocolRepository {
	return p.protocols
}
func (p *AgronomyProvider) Diseases() agronomyrepo.DiseaseRepository {
	return p.diseases
}
func (p *AgronomyProvider) StressProfiles() agronomyrepo.StressRepository {
	return p.stress
}

//
// OPERATIONS
//

type OperationsProvider struct {
	tasks      operationsrepo.TaskRepository
	operations operationsrepo.OperationRepository
}

func NewOperationsProvider() *OperationsProvider {
	return &OperationsProvider{
		tasks:      operationsinfra.NewTaskRepository(),
		operations: operationsinfra.NewOperationRepository(),
	}
}

func (p *OperationsProvider) Tasks() operationsrepo.TaskRepository {
	return p.tasks
}
func (p *OperationsProvider) Operations() operationsrepo.OperationRepository {
	return p.operations
}

//
// INVENTORY
//

type InventoryProvider struct {
	items     inventoryrepo.InventoryRepository
	movements inventoryrepo.MovementRepository
}

func NewInventoryProvider() *InventoryProvider {
	return &InventoryProvider{
		items:     inventoryinfra.NewInventoryRepository(),
		movements: inventoryinfra.NewMovementRepository(),
	}
}

func (p *InventoryProvider) Items() inventoryrepo.InventoryRepository {
	return p.items
}
func (p *InventoryProvider) Movements() inventoryrepo.MovementRepository {
	return p.movements
}

//
// ENVIRONMENT
//

type EnvironmentProvider struct {
	sensors      environmentrepo.SensorRepository
	telemetry    environmentrepo.TelemetryRepository
	climateZones environmentrepo.ClimateZoneRepository
}

func NewEnvironmentProvider() *EnvironmentProvider {
	return &EnvironmentProvider{
		sensors:      environmentinfra.NewSensorRepository(),
		telemetry:    environmentinfra.NewTelemetryRepository(),
		climateZones: environmentinfra.NewClimateZoneRepository(),
	}
}

func (p *EnvironmentProvider) Sensors() environmentrepo.SensorRepository {
	return p.sensors
}
func (p *EnvironmentProvider) Telemetry() environmentrepo.TelemetryRepository {
	return p.telemetry
}
func (p *EnvironmentProvider) ClimateZones() environmentrepo.ClimateZoneRepository {
	return p.climateZones
}

//
// AUTOMATION
//

type AutomationProvider struct {
	rules     automationrepo.RuleRepository
	actuators automationrepo.ActuatorRepository
}

func NewAutomationProvider() *AutomationProvider {
	return &AutomationProvider{
		rules:     automationinfra.NewRuleRepository(),
		actuators: automationinfra.NewActuatorRepository(),
	}
}

func (p *AutomationProvider) Rules() automationrepo.RuleRepository {
	return p.rules
}
func (p *AutomationProvider) Actuators() automationrepo.ActuatorRepository {
	return p.actuators
}

//
// ANALYTICS
//

type AnalyticsProvider struct {
	yieldRecords analyticsrepo.YieldRecordRepository
	metrics      analyticsrepo.MetricsRepository
}

func NewAnalyticsProvider() *AnalyticsProvider {
	return &AnalyticsProvider{
		yieldRecords: analyticsinfra.NewYieldRecordRepository(),

		metrics: analyticsinfra.NewMetricsRepository(),
	}
}

func (p *AnalyticsProvider) YieldRecords() analyticsrepo.YieldRecordRepository {
	return p.yieldRecords
}
func (p *AnalyticsProvider) Metrics() analyticsrepo.MetricsRepository {
	return p.metrics
}
