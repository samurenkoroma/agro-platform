package cropprotocol

type NutritionProfile struct {
	PHMin       *float64
	PHOptimal   *float64
	PHMax       *float64
	ECMin       *float64
	ECOptimal   *float64
	ECMax       *float64
	NTargetPPM  *float64
	PTargetPPM  *float64
	KTargetPPM  *float64
	CaTargetPPM *float64
	MgTargetPPM *float64
}
