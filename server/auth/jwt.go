package auth

import (
	"time"

	"github.com/6ar8nas/learning-go/server/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const TokenExpiryDuration = time.Hour * 1

func GenerateToken(userId uuid.UUID, isAdmin bool, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{}
	claims[types.ClaimsKeyUserId] = userId
	claims[types.ClaimsKeyIsAdmin] = isAdmin
	claims["exp"] = time.Now().Add(TokenExpiryDuration).Unix()

	return signString(claims, secretKey)
}

func VerifyToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, types.ErrorInvalidCredentials
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, types.ErrorInvalidCredentials
	}
}

func RefreshToken(tokenString string, secretKey []byte) (string, error) {
	claims, err := VerifyToken(tokenString, secretKey)
	if err != nil {
		return "", err
	}

	claims["exp"] = time.Now().Add(TokenExpiryDuration).Unix()
	return signString(claims, secretKey)
}

func signString(claims jwt.Claims, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
