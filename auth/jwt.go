package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const BearerSchema = "Bearer"
const tokenExpiryDuration = time.Hour * 1

func GenerateToken(claims map[string]interface{}, secretKey []byte) (string, error) {
	jwtClaims := jwt.MapClaims(claims)
	jwtClaims["exp"] = time.Now().Add(tokenExpiryDuration).Unix()

	return signString(jwtClaims, secretKey)
}

func VerifyToken(tokenString string, secretKey []byte) (map[string]interface{}, error) {
	return verifyToken(tokenString, secretKey)
}

func verifyToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorInvalidCredentials
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, ErrorInvalidCredentials
	}
}

func RefreshToken(tokenString string, secretKey []byte) (string, error) {
	claims, err := verifyToken(tokenString, secretKey)
	if err != nil {
		return "", err
	}

	claims["exp"] = time.Now().Add(tokenExpiryDuration).Unix()
	return signString(claims, secretKey)
}

func signString(claims jwt.Claims, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
