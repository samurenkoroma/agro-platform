package valueobject

type Metadata map[string]any

func NewMetadata() Metadata {
	return make(
		map[string]any,
	)
}
