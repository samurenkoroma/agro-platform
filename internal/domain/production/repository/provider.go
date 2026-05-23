package repository

type ProductionProvider interface {
	GrowingCycles() GrowingCycleRepository
	Plants() PlantRepository
	Slots() SlotRepository
	Substrates() SubstrateRepository
	Harvests() HarvestBatchRepository
	Yields() YieldBatchRepository
}
