package types

const HeaderXRequestId = "X-Request-Id"

type ContextKey string

const ContextKeyRequestId ContextKey = "request_id"
const ContextKeyUserId ContextKey = "user_id"
const ContextKeyIsAdmin ContextKey = "is_admin"

const ClaimsKeyUserId string = "user_id"
const ClaimsKeyIsAdmin string = "is_admin"
