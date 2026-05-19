package repository

import (
	"context"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
	ls "github.com/samurenkoroma/agro-platform/internal/domain/spatial/entity/layout_snapshot"
)

type LayoutSnapshotRepository interface {
	Save(ctx context.Context, snapshot *ls.Aggregate) error
	Get(ctx context.Context, id vo.ID) (*ls.Aggregate, error)
	GetLatest(ctx context.Context, farmID vo.ID) (*ls.Aggregate, error)
}
