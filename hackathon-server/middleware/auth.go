package middleware

import (
	"fmt"
	"hackathon-server/models"
	"hackathon-server/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(jwtSecret []byte) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.SendJSONResponse(w, http.StatusUnauthorized, models.Response{
					Success: false,
					Message: "Authorization header required",
				})
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				utils.SendJSONResponse(w, http.StatusUnauthorized, models.Response{
					Success: false,
					Message: "Invalid authorization format",
				})
				return
			}

			claims, err := utils.ValidateToken(tokenString, jwtSecret)
			if err != nil {
				utils.SendJSONResponse(w, http.StatusUnauthorized, models.Response{
					Success: false,
					Message: "Invalid or expired token",
				})
				return
			}

			r.Header.Set("X-User-ID", fmt.Sprintf("%d", claims.UserID))
			r.Header.Set("X-Username", claims.Username)

			next(w, r)
		}
	}
}