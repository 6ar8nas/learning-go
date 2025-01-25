package utils

import (
	"context"

	"github.com/6ar8nas/learning-go/types"
)

func GetContextValue(ctx context.Context, key types.ContextKey) any {
	return ctx.Value(key)
}

func AssignContextValue(ctx context.Context, key types.ContextKey, value any) context.Context {
	return context.WithValue(ctx, key, value)
}
