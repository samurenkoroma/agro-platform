package sensor

type Type string

const (
	Temperature Type = "TEMPERATURE"

	Humidity Type = "HUMIDITY"

	CO2 Type = "CO2"

	Light Type = "LIGHT"

	SoilMoisture Type = "SOIL_MOISTURE"

	PH Type = "PH"

	EC Type = "EC"

	WaterLevel Type = "WATER_LEVEL"

	FlowRate Type = "FLOW_RATE"

	Airflow Type = "AIRFLOW"
)
