package slot

type Capacity struct {
	MaxPlants int

	CurrentPlants int
}

func (c Capacity) HasSpace() bool {
	return c.CurrentPlants <
		c.MaxPlants
}

func (c Capacity) IsFull() bool {
	return c.CurrentPlants >=
		c.MaxPlants
}
