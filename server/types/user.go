package types

import sharedTypes "github.com/6ar8nas/learning-go/shared/types"

type UserHashedAuthRequest struct {
	Username string                     `json:"username"`
	Password sharedTypes.HashedPassword `json:"password"`
}
