package middleware

import "context"

const HeaderXRequestId = "X-Request-Id"

type ContextKey string

const ContextKeyRequestId ContextKey = "request_id"

func GetRequestId(ctx context.Context) string {
	return ctx.Value(ContextKeyRequestId).(string)
}
