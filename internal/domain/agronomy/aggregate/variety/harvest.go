package variety

type HarvestMode string

const (
	SingleHarvest     HarvestMode = "SINGLE"
	MultiHarvest      HarvestMode = "MULTI"
	ContinuousHarvest HarvestMode = "CONTINUOUS"
)

type HarvestProfile struct {
	Mode                 HarvestMode
	ExpectedHarvestCount *int
	HarvestWindowDays    *int
}
