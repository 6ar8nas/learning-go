package types

import "errors"

var ErrorNotFound = errors.New("requested resource was not found")
var ErrorPermissionDenied = errors.New("permission denied")
var ErrorInvalidCredentials = errors.New("invalid credentials")
var ErrorAuthenticationHeaderMissing = errors.New("authentication header missing or not sufficient")
