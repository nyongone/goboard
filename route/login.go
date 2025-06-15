package route

import (
	"database/sql"
	"go-board-api/handler"
	"go-board-api/repository"
	"go-board-api/service"

	"github.com/gorilla/mux"
)

func NewLoginRoute(api *mux.Router, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
  loginHandler := &handler.LoginHandler{Service: loginService}

  api.HandleFunc("/login", loginHandler.Login).Methods("POST")
}