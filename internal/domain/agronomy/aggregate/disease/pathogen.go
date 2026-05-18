package disease

type PathogenType string

const (
	Fungal        PathogenType = "FUNGAL"
	Bacterial     PathogenType = "BACTERIAL"
	Viral         PathogenType = "VIRAL"
	Oomycete      PathogenType = "OOMYCETE"
	Physiological PathogenType = "PHYSIOLOGICAL"
	Nutritional   PathogenType = "NUTRITIONAL"
	Unknown       PathogenType = "UNKNOWN"
)
