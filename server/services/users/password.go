package users

import (
	sharedTypes "github.com/6ar8nas/learning-go/shared/types"
	"golang.org/x/crypto/bcrypt"
)

const PasswordCost = 14

func HashPassword(password sharedTypes.Password) (sharedTypes.HashedPassword, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	return sharedTypes.HashedPassword(bytes), err
}

func VerifyPassword(password sharedTypes.Password, hash sharedTypes.HashedPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
