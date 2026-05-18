package cropprotocol

type ClimateProfile struct {
	TemperatureDayMin   *float64
	TemperatureDayMax   *float64
	TemperatureNightMin *float64
	TemperatureNightMax *float64
	HumidityMin         *float64
	HumidityMax         *float64
	CO2Min              *float64
	CO2Optimal          *float64
}
