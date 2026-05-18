package climatezone

import "errors"

var (
	ErrInvalidTemperatureRange = errors.New("invalid temperature range")
	ErrInvalidHumidityRange    = errors.New("invalid humidity range")
	ErrInvalidCO2Range         = errors.New("invalid co2 range")
	ErrArchivedZone            = errors.New("climate zone archived")
)
