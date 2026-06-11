package crop

import (
	"context"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type ListFilter struct {
	Search   *string
	Category *string
	Archived *bool
}

type Detail struct {
	ID             vo.ID       `json:"id"`
	Name           string      `json:"name"`
	Category       string      `json:"category"`
	Family         string      `json:"family"`
	ScientificName *string     `json:"scientificName"`
	Description    *string     `json:"description"`
	Metadata       vo.Metadata `json:"metadata"`
	ImageUrl       *string     `json:"imageUrl"`
}
type ListItem struct {
	ID             vo.ID   `json:"id"`
	Name           string  `json:"name"`
	ScientificName *string `json:"scientificName"`
	Category       string  `json:"category"`
	Family         string  `json:"family"`
	ImageUrl       *string `json:"imageUrl"`
}

type Projection interface {
	Get(ctx context.Context, Id string) (*Detail, error)
	List(ctx context.Context, filter ListFilter) ([]ListItem, error)
}
