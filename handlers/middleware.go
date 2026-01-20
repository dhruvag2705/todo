package handlers

import (
	"context"		// For context handling (user ID storage)
	"net/http"		//HTTP server and request handling
	"strings"		// For string manipulation (token separation for JWT)

	"github.com/golang-jwt/jwt/v5"	// JWT handling
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {		// Takes one API function and returns another
	return func(w http.ResponseWriter, r *http.Request) {			// The returned function
		tokenStr := r.Header.Get("Authorization")					// Get token from header
		if tokenStr == "" {
			http.Error(w, "No token", 403)							// if condition for No token found
			return
		}
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)		// if found, Remove "Bearer " prefix
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
