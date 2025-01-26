package middleware

import (
	"net/http"
	"strings"

	"github.com/6ar8nas/learning-go/auth"
	"github.com/6ar8nas/learning-go/server/config"
	"github.com/6ar8nas/learning-go/server/types"
	sharedTypes "github.com/6ar8nas/learning-go/shared/types"
	sharedUtils "github.com/6ar8nas/learning-go/shared/utils"
	"github.com/google/uuid"
)

var secretKey = config.AuthSecret

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/login") || strings.Contains(r.URL.Path, "/register") {
			next.ServeHTTP(w, r)
			return
		}

		tokenString, err := getToken(r)
		if err != nil {
			sharedUtils.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		claims, err := auth.VerifyToken(tokenString, secretKey)
		if err != nil {
			sharedUtils.WriteErrorJSON(w, http.StatusForbidden, err.Error())
			return
		}

		userId := uuid.MustParse(claims[types.ClaimsKeyUserId].(string))
		isAdmin := claims[types.ClaimsKeyIsAdmin].(bool)
		if !isAdmin && strings.Contains(r.URL.Path, "/users") {
			sharedUtils.WriteErrorJSON(w, http.StatusForbidden, sharedTypes.ErrorPermissionDenied.Error())
			return
		}

		ctx := sharedUtils.AssignContextValue(r.Context(), types.ContextKeyUserId, userId)
		ctx = sharedUtils.AssignContextValue(ctx, types.ContextKeyIsAdmin, isAdmin)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getToken(r *http.Request) (string, error) {
	tokenAuth := r.Header.Get("Authorization")

	if len(tokenAuth) > 40 { // We can safely assume token will be at least
		return tokenAuth[(len(auth.BearerSchema) + 1):], nil
	}

	return "", sharedTypes.ErrorAuthenticationHeaderMissing
}
