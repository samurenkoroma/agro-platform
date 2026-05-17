package shared

// Общие Value Objects, Events, Errors, Geometry и т.д.
type ID string

func (id ID) String() string {
	return string(id)
}
