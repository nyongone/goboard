package route

import (
	"database/sql"
	"go-board-api/handler"
	"go-board-api/repository"
	"go-board-api/service"

	"github.com/gorilla/mux"
)

func NewSignupRoute(api *mux.Router, db *sql.DB) {
  userRepository := repository.NewUserRepository(db)
  signupService := service.NewSignupService(userRepository)
  signupHandler := &handler.SignupHandler{Service: signupService}

  api.HandleFunc("/signup", signupHandler.Signup).Methods("POST")
}