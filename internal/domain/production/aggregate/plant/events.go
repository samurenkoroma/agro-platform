package plant

import (
	ev "github.com/samurenkoroma/agro-platform/internal/domain/shared/event"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

const (
	EventPlantCreated = "plant.created"

	EventPlantTransplanted = "plant.transplanted"

	EventPlantStageChanged = "plant.stage.changed"

	EventPlantStressed = "plant.stressed"

	EventPlantDiseased = "plant.diseased"

	EventPlantDiscarded = "plant.discarded"

	EventPlantHarvested = "plant.harvested"

	EventPlantDied = "plant.died"
)

type PlantCreated struct {
	ev.BaseEvent
}

func NewPlantCreated(
	id vo.ID,
) PlantCreated {

	return PlantCreated{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPlantCreated,
		),
	}
}

type PlantTransplanted struct {
	ev.BaseEvent

	ProductionUnitID vo.ID
}

func NewPlantTransplanted(
	id vo.ID,

	unitID vo.ID,
) PlantTransplanted {

	return PlantTransplanted{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPlantTransplanted,
		),

		ProductionUnitID: unitID,
	}
}

type PlantStageChanged struct {
	ev.BaseEvent

	StageID vo.ID
}

func NewPlantStageChanged(
	id vo.ID,

	stageID vo.ID,
) PlantStageChanged {

	return PlantStageChanged{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPlantStageChanged,
		),

		StageID: stageID,
	}
}

type PlantStressed struct {
	ev.BaseEvent
}

func NewPlantStressed(
	id vo.ID,
) PlantStressed {

	return PlantStressed{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPlantStressed,
		),
	}
}

type PlantDiseased struct {
	ev.BaseEvent
}

func NewPlantDiseased(
	id vo.ID,
) PlantDiseased {

	return PlantDiseased{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPlantDiseased,
		),
	}
}

type PlantDiscarded struct {
	ev.BaseEvent
}

func NewPlantDiscarded(
	id vo.ID,
) PlantDiscarded {

	return PlantDiscarded{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPlantDiscarded,
		),
	}
}

type PlantHarvested struct {
	ev.BaseEvent
}

func NewPlantHarvested(
	id vo.ID,
) PlantHarvested {

	return PlantHarvested{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPlantHarvested,
		),
	}
}

type PlantDied struct {
	ev.BaseEvent
}

func NewPlantDied(
	id vo.ID,
) PlantDied {

	return PlantDied{
		BaseEvent: ev.NewBaseEvent(
			id,
			EventPlantDied,
		),
	}
}
