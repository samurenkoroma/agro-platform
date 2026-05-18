package movement

type ReferenceType string

const (
	TaskReference          ReferenceType = "TASK"
	GrowingCycleReference  ReferenceType = "GROWING_CYCLE"
	HarvestReference       ReferenceType = "HARVEST"
	FertilizationReference ReferenceType = "FERTILIZATION"
	ManualReference        ReferenceType = "MANUAL"
	CorrectionReference    ReferenceType = "CORRECTION"
)

type Reference struct {
	Type ReferenceType
	ID   string
}
