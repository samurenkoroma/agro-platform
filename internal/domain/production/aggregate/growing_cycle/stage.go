package growingcycle

type CycleStage string

const (
	StagePlanning    CycleStage = "planning"
	StageGermination CycleStage = "germination"
	StageSeedling    CycleStage = "seedling"
	StageVegetative  CycleStage = "vegetative"
	StageFlowering   CycleStage = "flowering"
	StageFruiting    CycleStage = "fruiting"
	StageHarvesting  CycleStage = "harvesting"
	StageCompleted   CycleStage = "completed"
)
