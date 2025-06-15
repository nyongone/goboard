package route

import (
	"database/sql"
	"go-board-api/handler"
	"go-board-api/middleware"
	"go-board-api/repository"
	"go-board-api/service"

	"github.com/gorilla/mux"
)

func NewUserRoute(api *mux.Router, db *sql.DB) {
  userRepository := repository.NewUserRepository(db)
  userService := service.NewUserService(userRepository)
  userHandler := &handler.UserHandler{Service: userService}

  api.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
  api.HandleFunc("/users/{id:[0-9]+}", middleware.AuthMiddleware(db, userHandler.UpdateUser)).Methods("PUT")
}