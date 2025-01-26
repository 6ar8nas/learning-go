package auth

import "errors"

var ErrorInvalidCredentials = errors.New("invalid credentials")
var ErrorTokenIsExpired = errors.New("expired token")
