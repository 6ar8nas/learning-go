package utils

import (
	"6ar8nas/test-app/types"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password types.Password) (types.HashedPassword, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return types.HashedPassword(bytes), err
}

func VerifyPassword(password types.Password, hash types.HashedPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
