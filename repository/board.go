package repository

import (
	"database/sql"
	"go-board-api/model"
)

type boardRepository struct {
  db *sql.DB
}

func NewBoardRepository(db *sql.DB) model.BoardRepository {
  return &boardRepository{db}
}

func (br *boardRepository) FindAll() ([]*model.Board, error) {
  rows, err := br.db.Query("SELECT id, name, slug, created_at, updated_at, deleted_at FROM boards ORDER BY created_at DESC")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var boards []*model.Board
  for rows.Next() {
    board := &model.Board{}
    err := rows.Scan(&board.ID, &board.Name, &board.Slug, &board.CreatedAt, &board.UpdatedAt, &board.DeletedAt)
    if err != nil {
      return nil, err
    }

    boards = append(boards, board)
  }

  return boards, nil
}