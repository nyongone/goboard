package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-board-api/config"
	"go-board-api/datastore"
	"go-board-api/handler"
	"go-board-api/internal/logger"
	"go-board-api/middleware"
	"go-board-api/model"
	"go-board-api/repository"
	"go-board-api/service"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

func main() {
  config.LoadEnv()
  logger.Init()

  db, err := datastore.OpenConnection()
  if err != nil {
    logger.Error("An error occured while connecting mysql database", zap.Error(err))
  }

  // Set max idle/open connections
  idleConns, _ := strconv.Atoi(config.EnvVar.DBMaxIdleConns)
  maxOpenConns, _ := strconv.Atoi(config.EnvVar.DBMaxOpenConns)

  db.SetMaxIdleConns(idleConns)
  db.SetMaxOpenConns(maxOpenConns)

  err = db.Ping()
  if err != nil {
    logger.Error("An error occured while connecting mysql database", zap.Error(err))
  }

  r := mux.NewRouter()
  api := r.PathPrefix("/api").Subrouter()
  v1 := api.PathPrefix("/v1").Subrouter()

  r.Use(middleware.Logger)

  r.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(&model.Response{
      Code: http.StatusOK,
      Message: "OK",
    })
  })

  userRepository := repository.NewUserRepository(db)
  userService := service.NewUserService(userRepository)
  userHandler := &handler.UserHandler{Service: userService}

  postRepository := repository.NewPostRepository(db)
  postService := service.NewPostService(postRepository)
  postHandler := &handler.PostHandler{Service: postService}

  commentRepository := repository.NewCommentRepository(db)
  commentService := service.NewCommentService(commentRepository)
  commentHandler := &handler.CommentHandler{Service: commentService}

  boardRepository := repository.NewBoardRepository(db)
  boardService := service.NewBoardService(boardRepository)
  boardHandler := &handler.BoardHandler{Service: boardService}

  signupService := service.NewSignupService(userRepository)
  signupHandler := &handler.SignupHandler{Service: signupService}

  loginService := service.NewLoginService(userRepository)
  loginHandler := &handler.LoginHandler{Service: loginService}

  // User Handlers
  v1.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
  v1.HandleFunc("/users/{id:[0-9]+}", middleware.AuthMiddleware(db, userHandler.UpdateUser)).Methods("PUT")

  // Post Handlers
  v1.HandleFunc("/posts/{id:[0-9]+}", postHandler.GetPostById).Methods("GET")
  v1.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET")
  v1.HandleFunc("/posts/byBoard/{board_id:[0-9]+}", postHandler.GetAllPostsByBoard).Methods("GET")
  v1.HandleFunc("/posts/byAuthor/{author_id:[0-9]+}", postHandler.GetAllPostsByAuthor).Methods("GET")
  v1.HandleFunc("/posts", middleware.AuthMiddleware(db, postHandler.WritePost)).Methods("POST")
  v1.HandleFunc("/posts/{id:[0-9]+}", middleware.AuthMiddleware(db, postHandler.UpdatePost)).Methods("PUT")
  v1.HandleFunc("/posts/{id:[0-9]+}", middleware.AuthMiddleware(db, postHandler.DeletePost)).Methods("DELETE")

  // Comment Handlers
  v1.HandleFunc("/comments/byPost/{id:[0-9]+}", commentHandler.GetCommentsByPost).Methods("GET")
  v1.HandleFunc("/comments/byAuthor/{id:[0-9]+}", commentHandler.GetCommentByAuthor).Methods("GET")
  v1.HandleFunc("/comments", middleware.AuthMiddleware(db, commentHandler.WriteComment)).Methods("POST")
  v1.HandleFunc("/comments/{id:[0-9]+}", middleware.AuthMiddleware(db, commentHandler.DeleteComment)).Methods("DELETE")

  // Board Handlers
  v1.HandleFunc("/boards", boardHandler.GetAllBoards).Methods("GET")

  // Signup Handlers
  v1.HandleFunc("/signup", signupHandler.Signup).Methods("POST")

  // Login Handlers
  v1.HandleFunc("/login", loginHandler.Login).Methods("POST")

  c := cors.New(cors.Options{
    AllowedHeaders: strings.Split(config.EnvVar.AppCorsHeaders, ","),
    AllowedMethods: strings.Split(config.EnvVar.AppCorsMethods, ","),
    AllowedOrigins: strings.Split(config.EnvVar.AppCorsOrigins, ","),
    AllowCredentials: true,
  })

  handler := c.Handler(r)

  // http.ListenAndServe(fmt.Sprintf("%s:%s", config.EnvVar.AppHost, config.EnvVar.AppPort), handler)

  svr := &http.Server{
    Addr: fmt.Sprintf("%s:%s", config.EnvVar.AppHost, config.EnvVar.AppPort),
    Handler: handler,
  }

  go func() {
    logger.Info("Server running...", 
                zap.String("host", config.EnvVar.AppHost),
                zap.String("port", config.EnvVar.AppPort),
                zap.String("cors_headers", config.EnvVar.AppCorsHeaders),
                zap.String("cors_methods", config.EnvVar.AppCorsMethods),
                zap.String("cors_origins", config.EnvVar.AppCorsOrigins))

    if svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      logger.Fatal("An error occured while running server", zap.Error(err))
    }
  }()

  stop := GracefulShutdown(func() {
    logger.Info("Shutting down Server")
    if err := svr.Shutdown(context.Background()); err != nil {
      logger.Error("An error occured while shutting down server", zap.Error(err))
    }

    logger.Info("Close DB Connection")
    if err := db.Close(); err != nil {
      logger.Error("An error occured while closing db connection", zap.Error(err))
    }

    logger.Info("Gracefully Stopped Server")
  }, syscall.SIGINT, syscall.SIGTERM)

  <-stop
}

func GracefulShutdown(fn func(), sigs ...os.Signal) <-chan struct{} {
  stop := make(chan struct{})
  sigChan := make(chan os.Signal, 1)

  signal.Notify(sigChan, sigs...)

  go func() {
    <-sigChan
    signal.Stop(sigChan)
    
    fn()

    close(sigChan)
    close(stop)
  }()

  return stop
}