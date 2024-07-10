package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/akposieyefa/open-ai/config"
	"github.com/akposieyefa/open-ai/core/models"
	response "github.com/akposieyefa/open-ai/pkg/http"
	"github.com/golang-jwt/jwt"
)

// jwt verify
func JwtVerify(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization") //Grab the token from the header

		tokenString = strings.TrimSpace(tokenString)

		if tokenString == "" {
			response.WriteErrorJson("Unauthorized", "Missing auth token", false, w, http.StatusUnauthorized)
			return
		}

		tk := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, tk, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Envs.JWT_SECRET), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				response.WriteErrorJson("Invalid auth signature", err.Error(), false, w, http.StatusUnauthorized)
				return
			}
			response.WriteErrorJson("Error parsing token", err.Error(), false, w, http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			response.WriteErrorJson("Token is not valid", "invalid token", false, w, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
