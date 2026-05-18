package variety

type GrowthProfile struct {
	MinHeightCM      *float64
	MaxHeightCM      *float64
	RootDepthCM      *float64
	CanopyDiameterCM *float64
	Determinate      bool
	SupportsPruning  bool
	SupportsTrellis  bool
}
