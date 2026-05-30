package productionunit

import vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"

type Properties struct {
	Capacity       Capacity       `json:"capacity"`
	ClimateProfile ClimateProfile `json:"climateProfile"`
	Capabilities   []Capability   `json:"capabilities"`
	Metadata       vo.Metadata    `json:"metadata"`
	Position       *vo.Position   `json:"position"`
}

func (p *Properties) AddCapabilities(capabilities []string) {
	for _, capability := range capabilities {
		p.Capabilities = append(p.Capabilities, Capability(capability))
	}
}

func NewProps(name, desc string) *Properties {
	meta := vo.NewMetadata()
	meta["name"] = name
	meta["description"] = desc
	return &Properties{
		Capacity:       Capacity{},
		ClimateProfile: ClimateProfile{},
		Capabilities:   []Capability{},
		Metadata:       meta,
		Position: &vo.Position{
			X: 0,
			Y: 0,
			Z: nil,
		},
	}
}

type ClimateProfile struct {
	TemperatureMin *float64 `json:"temperatureMin"`
	TemperatureMax *float64 `json:"temperatureMax"`

	HumidityMin *float64 `json:"humidityMin"`
	HumidityMax *float64 `json:"humidityMax"`

	CO2Min *float64 `json:"co2Min"`
	CO2Max *float64 `json:"co2Max"`

	LightPPFDMin *float64 `json:"lightPPFDMin"`
	LightPPFDMax *float64 `json:"lightPPFDMax"`
}
type Capacity struct {
	PlantCapacity     *int     `json:"plantCapacity"`
	WaterVolumeLiters *float64 `json:"waterVolumeLiters"`
	AreaM2            *float64 `json:"areaM2"`
	TrayCount         *int     `json:"trayCount"`
	ChannelCount      *int     `json:"channelCount"`
}

type Capability string

const (
	// soil
	Soil Capability = "SOIL"

	// water
	Irrigation  Capability = "IRRIGATION"
	Fertigation Capability = "FERTIGATION"
	Drainage    Capability = "DRAINAGE"

	// hydro
	Hydroponic      Capability = "HYDROPONIC"
	Aeroponic       Capability = "AEROPONIC"
	NutrientControl Capability = "NUTRIENT_CONTROL"

	// climate
	Lighting       Capability = "LIGHTING"
	ClimateControl Capability = "CLIMATE_CONTROL"

	// iot
	SensorSupport Capability = "SENSOR_SUPPORT"
	Automation    Capability = "AUTOMATION"
	SlotBased     Capability = "SLOT_BASED"
	Mobile        Capability = "MOBILE"
)
