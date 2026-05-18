package stress

type StressType string

const (
	HeatStress         StressType = "HEAT"
	ColdStress         StressType = "COLD"
	DroughtStress      StressType = "DROUGHT"
	WaterloggingStress StressType = "WATERLOGGING"
	LightExcessStress  StressType = "LIGHT_EXCESS"
	LightDeficitStress StressType = "LIGHT_DEFICIT"
	SalinityStress     StressType = "SALINITY"
	PHStress           StressType = "PH"
	ECStress           StressType = "EC"
	TransplantStress   StressType = "TRANSPLANT"
	NutrientStress     StressType = "NUTRIENT"
	MechanicalStress   StressType = "MECHANICAL"
)
