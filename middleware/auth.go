package middleware

import (
	"6ar8nas/test-app/auth"
	"6ar8nas/test-app/config"
	"6ar8nas/test-app/types"
	"6ar8nas/test-app/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

var secretKey = config.AuthSecret

const BEARER_SCHEMA = "Bearer"

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/login") || strings.Contains(r.URL.Path, "/register") {
			next.ServeHTTP(w, r)
			return
		}

		tokenString, err := getToken(r)
		if err != nil {
			utils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		claims, err := auth.VerifyToken(tokenString, secretKey)
		if err != nil {
			utils.WriteErrorJSON(w, http.StatusForbidden, err.Error())
			return
		}

		userId := uuid.MustParse(claims[types.ClaimsKeyUserId].(string))
		isAdmin := claims[types.ClaimsKeyIsAdmin].(bool)
		if !isAdmin && strings.Contains(r.URL.Path, "/users") {
			utils.WriteErrorJSON(w, http.StatusForbidden, types.ErrorPermissionDenied.Error())
			return
		}

		ctx := utils.AssignContextValue(r.Context(), types.ContextKeyUserId, userId)
		ctx = utils.AssignContextValue(ctx, types.ContextKeyIsAdmin, isAdmin)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getToken(r *http.Request) (string, error) {
	tokenAuth := r.Header.Get("Authorization")

	if len(tokenAuth) > 40 { // We can safely assume token will be at least
		return tokenAuth[(len(BEARER_SCHEMA) + 1):], nil
	}

	return "", types.ErrorAuthenticationHeaderMissing
}
