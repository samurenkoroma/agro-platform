package getproductionunit

type Result struct {
	ID string `json:"id"`

	FarmID string `json:"farmId"`

	Name string `json:"name"`

	Type string `json:"type"`

	ParentID *string `json:"parentId"`

	CreatedAt string `json:"createdAt"`

	UpdatedAt string `json:"updatedAt"`
}
