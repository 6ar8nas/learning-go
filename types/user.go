package types

import "github.com/google/uuid"

type HashedPassword string

type User struct {
	Id       uuid.UUID      `json:"id"`
	Username string         `json:"username"`
	Admin    bool           `json:"-"`
	Password HashedPassword `json:"-"`
}

type Password string

type UserAuthRequest struct {
	Username string   `json:"username"`
	Password Password `json:"password"`
}

type UserHashedAuthRequest struct {
	Username string         `json:"username"`
	Password HashedPassword `json:"password"`
}

type UserAuthResponse struct {
	AuthToken string `json:"auth_token"` // TODO: JWT
}
