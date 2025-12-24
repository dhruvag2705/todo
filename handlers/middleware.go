package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "No token", 403)
			return
		}
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", 401)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		uid := int(claims["userId"].(float64))

		ctx := context.WithValue(r.Context(), "userId", uid)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
