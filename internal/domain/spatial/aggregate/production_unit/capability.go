package productionunit

type Capability string

const (

	// soil

	Soil Capability = "SOIL"

	// water

	Irrigation Capability = "IRRIGATION"

	Fertigation Capability = "FERTIGATION"

	Drainage Capability = "DRAINAGE"

	// hydro

	Hydroponic Capability = "HYDROPONIC"

	Aeroponic Capability = "AEROPONIC"

	NutrientControl Capability = "NUTRIENT_CONTROL"

	// climate

	Lighting Capability = "LIGHTING"

	ClimateControl Capability = "CLIMATE_CONTROL"

	// iot

	SensorSupport Capability = "SENSOR_SUPPORT"

	Automation Capability = "AUTOMATION"

	SlotBased Capability = "SLOT_BASED"

	Mobile Capability = "MOBILE"
)
