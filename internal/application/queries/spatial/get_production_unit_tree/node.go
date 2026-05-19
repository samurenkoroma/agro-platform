package getproductionunittree

type Node struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`

	Children []Node `json:"children"`
}
