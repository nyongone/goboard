package service

import (
	"go-board-api/internal/util"
	"go-board-api/model"
)

type loginService struct {
	repo model.UserRepository
}

func NewLoginService(repo model.UserRepository) model.LoginService {
	return &loginService{repo}
}

func (ls *loginService) FindOneByEmail(email string) (*model.User, error) {
	return ls.repo.FindOneByEmail(email)
}

func (ls *loginService) CreateAccessToken(email string) (string, error) {
	return util.CreateAccessToken(email)
}

func (ls *loginService) CreateRefreshToken() (string, error) {
	return util.CreateRefreshToken()
}

func (ls *loginService) UpdateRefreshToken(id int, token string) error {
	return ls.repo.UpdateRefreshToken(id, token)
}