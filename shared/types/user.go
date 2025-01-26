package types

import "github.com/google/uuid"

type Password string
type HashedPassword string

type User struct {
	Id       uuid.UUID      `json:"id"`
	Username string         `json:"username"`
	Admin    bool           `json:"-"`
	Password HashedPassword `json:"-"`
}

type UserAuthRequest struct {
	Username string   `json:"username"`
	Password Password `json:"password"`
}

type UserAuthResponse struct {
	AuthToken string `json:"auth_token"`
}
