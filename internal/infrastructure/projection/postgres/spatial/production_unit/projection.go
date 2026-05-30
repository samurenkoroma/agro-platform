package productionunit

import (
	"context"
	"encoding/json"

	productionunit "github.com/samurenkoroma/agro-platform/internal/application/queries/spatial/production_unit"
	"github.com/samurenkoroma/agro-platform/internal/application/uow"
	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

type projection struct {
	db uow.DB
}

func New(db uow.DB) productionunit.Projection {
	return &projection{db: db}
}

func (p *projection) Get(ctx context.Context, id vo.ID) (*productionunit.DTO, error) {
	sql := `SELECT id,parent_id,type,status,name,code,description,geometry,position,capacity,climate,properties,metadata
FROM production_units WHERE id = $1 LIMIT 1`

	row := p.db.QueryRow(ctx, sql, id)

	return scanDTO(row)
}

func (p *projection) ListRoots(ctx context.Context, ownerId vo.ID) ([]productionunit.DTO, error) {
	sql := `SELECT id,parent_id,type,status,code,geometry,properties
FROM production_units 
WHERE owner_id = $1 AND parent_id IS NULL AND archived_at IS NULL 
ORDER BY code`

	rows, err := p.db.Query(ctx, sql, ownerId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]productionunit.DTO, 0)

	for rows.Next() {
		item, err := scanDTO(rows)

		if err != nil {
			return nil, err
		}

		result = append(result, *item)
	}

	return result, nil
}

func (p *projection) ListChildren(ctx context.Context, parentID vo.ID) ([]productionunit.DTO, error) {

	sql := `SELECT id,parent_id,type,status,name,code,description,geometry,position,capacity,climate,properties,metadata
FROM production_units
WHERE parent_id = $1 AND archived_at IS NULL ORDER BY name`

	rows, err := p.db.Query(
		ctx,
		sql,
		parentID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]productionunit.DTO, 0)

	for rows.Next() {
		item, err := scanDTO(rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *item)
	}

	return result, nil
}

func (p *projection) Tree(ctx context.Context, rootID *vo.ID) ([]productionunit.TreeNode, error) {
	sql := `WITH RECURSIVE tree AS (
SELECT id, parent_id, type, status, name
FROM production_units
WHERE (
		($1::uuid IS NULL AND parent_id IS NULL)
		OR
		(id = $1)
	)

	UNION ALL

	SELECT p.id, p.parent_id, p.type, p.status, p.name
	FROM production_units p
	INNER JOIN tree t
		ON p.parent_id = t.id
)
SELECT id,parent_id,type,status,name
FROM tree
ORDER BY name`

	rows, err := p.db.Query(ctx, sql, rootID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	type rawNode struct {
		ID       vo.ID
		ParentID *vo.ID
		Type     string
		Status   string
		Name     string
	}

	nodes := make([]rawNode, 0)

	for rows.Next() {

		var item rawNode

		err = rows.Scan(&item.ID, &item.ParentID, &item.Type, &item.Status, &item.Name)

		if err != nil {
			return nil, err
		}

		nodes = append(nodes, item)
	}

	index := make(
		map[vo.ID]*productionunit.TreeNode,
	)

	for _, node := range nodes {

		index[node.ID] = &productionunit.TreeNode{
			ID: node.ID,

			ParentID: node.ParentID,

			Type: node.Type,

			Status: node.Status,

			Name: node.Name,

			Children: make(
				[]productionunit.TreeNode,
				0,
			),
		}
	}

	result := make(
		[]productionunit.TreeNode,
		0,
	)

	for _, node := range nodes {

		current :=
			index[node.ID]

		if node.ParentID == nil {

			result = append(
				result,
				*current,
			)

			continue
		}

		parent, ok :=
			index[*node.ParentID]

		if !ok {
			continue
		}

		parent.Children = append(
			parent.Children,
			*current,
		)
	}

	return result, nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanDTO(
	row scanner,
) (*productionunit.DTO, error) {

	var result productionunit.DTO

	var geometryRaw []byte

	var propertiesRaw []byte

	err := row.Scan(
		&result.ID,
		&result.ParentID,
		&result.Type,
		&result.Status,
		&result.Code,
		&geometryRaw,
		&propertiesRaw,
	)

	if err != nil {
		return nil, err
	}

	if propertiesRaw != nil {
		var props map[string]any

		if err := json.Unmarshal(propertiesRaw, &props); err != nil {
			return nil, err
		}
		result.Properties = props
	}

	return &result, nil
}
