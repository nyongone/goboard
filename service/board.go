package service

import "go-board-api/model"

type boardService struct {
	repo model.BoardRepository
}

func NewBoardService(repo model.BoardRepository) model.BoardService {
	return &boardService{repo}
}

func (bs *boardService) FindAll() ([]*model.Board, error) {
	return bs.repo.FindAll()
}