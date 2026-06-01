package repository

type ProductionProvider interface {
	GrowingCycles() GrowingCycleRepository
	Harvests() HarvestBatchRepository
	Planting() PlantingRepository
	Allocation() AllocationRepository
}
