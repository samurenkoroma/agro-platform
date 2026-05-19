package getproductionunittree

import (
	"context"

	repository "github.com/samurenkoroma/agro-platform/internal/shared/repository"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type Handler struct {
	repositories repository.Provider
}

func New(repositories repository.Provider) *Handler {
	return &Handler{
		repositories: repositories,
	}
}

func (h *Handler) Handle(ctx context.Context, query Query) (Node, error) {
	root, err := h.repositories.
		Spatial().
		ProductionUnits().
		GetByID(
			query.RootID,
		)

	if err != nil {
		return Node{}, err
	}

	if root == nil {
		return Node{},
			ErrRootNotFound
	}

	return h.buildTree(
		root.ID,
	)
}

func (h *Handler) buildTree(id vo.ID) (Node, error) {

	unit,
		err :=
		h.repositories.
			Spatial().
			ProductionUnits().
			GetByID(
				id,
			)

	if err != nil {
		return Node{}, err
	}

	children,
		err :=
		h.repositories.
			Spatial().
			ProductionUnits().
			GetChildren(
				id,
			)

	if err != nil {
		return Node{}, err
	}

	node := Node{
		ID: unit.ID.String(),

		Name: unit.Name,

		Type: string(
			unit.Type,
		),

		Children: make(
			[]Node,
			0,
		),
	}

	for _, child := range children {

		childNode,
			err :=
			h.buildTree(
				child.ID,
			)

		if err != nil {
			return Node{}, err
		}

		node.Children =
			append(
				node.Children,
				childNode,
			)
	}

	return node,
		nil
}
