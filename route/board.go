package route

import (
	"database/sql"
	"go-board-api/handler"
	"go-board-api/repository"
	"go-board-api/service"

	"github.com/gorilla/mux"
)

func NewBoardRoute(api *mux.Router, db *sql.DB) {
  boardRepository := repository.NewBoardRepository(db)
  boardService := service.NewBoardService(boardRepository)
  boardHandler := &handler.BoardHandler{Service: boardService}

  api.HandleFunc("/boards", boardHandler.GetAllBoards).Methods("GET")
}