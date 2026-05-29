package productionunit

type ClimateProfile struct {
	TemperatureMin *float64
	TemperatureMax *float64

	HumidityMin *float64
	HumidityMax *float64

	CO2Min *float64
	CO2Max *float64

	LightPPFDMin *float64
	LightPPFDMax *float64
}
