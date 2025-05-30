package middleware

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-board-api/internal/util"
	"go-board-api/model"
	"go-board-api/repository"
	"go-board-api/service"
	"net/http"
	"strings"
)

func AuthMiddleware(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authHeader := strings.Split(r.Header.Get("Authorization"), " ")

		if len(authHeader) == 2 {
			access_token := authHeader[1]
			authorized, err := util.ValidateToken(access_token)
			if !authorized || err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&model.Response{
					Code: http.StatusUnauthorized,
					Message: "UNAUTHORIZED",	
				})

				return
			}

			decoded, err := util.DecodeToken(access_token)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&model.Response{
					Code: http.StatusUnauthorized,
					Message: "UNAUTHORIZED",	
				})

				return
			}

			user, err := userService.FindOneByEmail(decoded.Email)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&model.Response{
					Code: http.StatusUnauthorized,
					Message: "UNAUTHORIZED",	
				})

				return
			}

			r = r.WithContext(context.WithValue(r.Context(), model.UserKey, user))
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&model.Response{
				Code: http.StatusUnauthorized,
				Message: "UNAUTHORIZED",	
			}) 

			return
		}

	})
}