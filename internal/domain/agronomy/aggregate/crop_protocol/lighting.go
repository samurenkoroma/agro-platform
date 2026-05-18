package cropprotocol

type LightingProfile struct {
	PPFDMin                   *float64
	PPFDTarget                *float64
	PPFDMax                   *float64
	DLI                       *float64
	PhotoperiodHours          *float64
	SupportsSupplementalLight bool
}
