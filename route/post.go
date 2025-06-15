package route

import (
	"database/sql"
	"go-board-api/handler"
	"go-board-api/middleware"
	"go-board-api/repository"
	"go-board-api/service"

	"github.com/gorilla/mux"
)

func NewPostRoute(api *mux.Router, db *sql.DB) {
  postRepository := repository.NewPostRepository(db)
  postService := service.NewPostService(postRepository)
  postHandler := &handler.PostHandler{Service: postService}

  api.HandleFunc("/posts/{id:[0-9]+}", postHandler.GetPostById).Methods("GET")
  api.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET")
  api.HandleFunc("/posts/byBoard/{board_id:[0-9]+}", postHandler.GetAllPostsByBoard).Methods("GET")
  api.HandleFunc("/posts/byAuthor/{author_id:[0-9]+}", postHandler.GetAllPostsByAuthor).Methods("GET")
  api.HandleFunc("/posts", middleware.AuthMiddleware(db, postHandler.WritePost)).Methods("POST")
  api.HandleFunc("/posts/{id:[0-9]+}", middleware.AuthMiddleware(db, postHandler.UpdatePost)).Methods("PUT")
  api.HandleFunc("/posts/{id:[0-9]+}", middleware.AuthMiddleware(db, postHandler.DeletePost)).Methods("DELETE")
}