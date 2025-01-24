package utils

import (
	"6ar8nas/test-app/types"
	"context"
)

func GetContextValue(ctx context.Context, key string) string {
	return ctx.Value(key).(string)
}

func AssignContextValue(ctx context.Context, key types.ContextKey, value any) context.Context {
	return context.WithValue(ctx, key, value)
}
