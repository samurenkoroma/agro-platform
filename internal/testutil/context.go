package testutil

import (
	"context"

	vo "github.com/samurenkoroma/agro-platform/internal/domain/shared/valueobject"
)

// OrgCtx возвращает контекст с organization_id.
func OrgCtx() context.Context {
	return OrgCtxWithID(vo.NewID())
}

// OrgCtxWithID возвращает контекст с конкретным organization_id.
func OrgCtxWithID(id vo.ID) context.Context {
	return context.WithValue(context.Background(), "organization_id", id.String())
}

// UserCtx возвращает контекст с organization_id и user_id.
func UserCtx() context.Context {
	ctx := OrgCtx()
	return context.WithValue(ctx, "user_id", vo.NewID().String())
}
