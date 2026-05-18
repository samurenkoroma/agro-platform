package cropprotocol

type IrrigationProfile struct {
	IrrigationEventsPerDay *int
	RunDurationSeconds     *int
	DrainTargetPercent     *float64
	PulseMode              bool
	Recirculation          bool
}
