package utils

import (
	"6ar8nas/test-app/types"
	"context"
)

func GetContextValue(ctx context.Context, key types.ContextKey) any {
	return ctx.Value(key)
}

func AssignContextValue(ctx context.Context, key types.ContextKey, value any) context.Context {
	return context.WithValue(ctx, key, value)
}
