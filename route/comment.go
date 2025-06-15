package route

import (
	"database/sql"
	"go-board-api/handler"
	"go-board-api/middleware"
	"go-board-api/repository"
	"go-board-api/service"

	"github.com/gorilla/mux"
)

func NewCommentRoute(api *mux.Router, db *sql.DB) {
  commentRepository := repository.NewCommentRepository(db)
  commentService := service.NewCommentService(commentRepository)
  commentHandler := &handler.CommentHandler{Service: commentService}

  api.HandleFunc("/comments/byPost/{id:[0-9]+}", commentHandler.GetCommentsByPost).Methods("GET")
  api.HandleFunc("/comments/byAuthor/{id:[0-9]+}", commentHandler.GetCommentByAuthor).Methods("GET")
  api.HandleFunc("/comments", middleware.AuthMiddleware(db, commentHandler.WriteComment)).Methods("POST")
  api.HandleFunc("/comments/{id:[0-9]+}", middleware.AuthMiddleware(db, commentHandler.DeleteComment)).Methods("DELETE")
}