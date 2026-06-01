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
	sql := `SELECT id,parent_id,type,status,code,geometry,properties FROM production_units WHERE id = $1`

	row := p.db.QueryRow(ctx, sql, id)

	return scanDTO(row)
}

func (p *projection) ListRoots(ctx context.Context, ownerId vo.ID) ([]*productionunit.DTO, error) {
	sql := `SELECT id,parent_id,type,status,code,geometry,properties
FROM production_units 
WHERE owner_id = $1 AND parent_id IS NULL AND archived_at IS NULL 
ORDER BY code`

	rows, err := p.db.Query(ctx, sql, ownerId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*productionunit.DTO, 0)

	for rows.Next() {
		item, err := scanDTO(rows)

		if err != nil {
			return nil, err
		}
		dto, err := p.Tree(ctx, &item.ID)
		result = append(result, dto)
	}

	return result, nil
}

func (p *projection) Tree(ctx context.Context, rootID *vo.ID) (*productionunit.DTO, error) {

	sql := `
WITH RECURSIVE tree AS (
	SELECT id, parent_id, type, status, code, geometry, properties
	FROM production_units
	WHERE (
		($1::uuid IS NULL AND parent_id IS NULL)
		OR
		(id = $1)
	)

	UNION ALL

	SELECT p.id, p.parent_id, p.type, p.status, p.code, p.geometry, p.properties
	FROM production_units p
	INNER JOIN tree t
		ON p.parent_id = t.id
)
SELECT id, parent_id, type, status, code, geometry, properties
FROM tree
ORDER BY code
`

	rows, err := p.db.Query(ctx, sql, rootID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type rawNode struct {
		ID         vo.ID
		ParentID   *vo.ID
		Type       string
		Status     string
		Code       string
		Geometry   *vo.Geometry
		Properties map[string]any
	}

	var nodes []rawNode

	for rows.Next() {
		var item rawNode

		if err := rows.Scan(
			&item.ID,
			&item.ParentID,
			&item.Type,
			&item.Status,
			&item.Code,
			&item.Geometry,
			&item.Properties,
		); err != nil {
			return nil, err
		}

		nodes = append(nodes, item)
	}

	nodeMap := make(map[vo.ID]rawNode)
	childrenMap := make(map[vo.ID][]vo.ID)

	var roots []vo.ID

	for _, n := range nodes {
		nodeMap[n.ID] = n

		if n.ParentID == nil {
			roots = append(roots, n.ID)
			continue
		}

		childrenMap[*n.ParentID] = append(
			childrenMap[*n.ParentID],
			n.ID,
		)
	}

	var build func(id vo.ID) *productionunit.DTO

	build = func(id vo.ID) *productionunit.DTO {
		n := nodeMap[id]

		result := productionunit.DTO{
			ID:         n.ID,
			ParentID:   n.ParentID,
			Type:       n.Type,
			Status:     n.Status,
			Code:       n.Code,
			Geometry:   n.Geometry,
			Properties: n.Properties,
			Children:   make([]*productionunit.DTO, 0),
		}

		for _, childID := range childrenMap[id] {
			result.Children = append(result.Children, build(childID))
		}

		return &result
	}

	return build(roots[0]), nil
}

type scanner interface {
	Scan(dest ...any) error
}

func scanDTO(row scanner) (*productionunit.DTO, error) {

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
