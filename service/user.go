package service

import (
	"errors"
	"go-board-api/model"
)

type userService struct {
  repo model.UserRepository
}

func NewUserService(repo model.UserRepository) model.UserService {
  return &userService{repo}
}

func (us *userService) FindOne(id int) (*model.User, error) {
  return us.repo.FindOne(id)
}

func (us *userService) FindOneByEmail(email string) (*model.User, error) {
  return us.repo.FindOneByEmail(email)
}

func (us *userService) Create(user *model.User) error {
  if _, err := us.repo.FindOneByEmail(user.Email); err == nil {
    return errors.New("user already exists")
  }

  return us.repo.Create(user)
}

func (us *userService) Update(id int, user *model.User) error {
  if _, err := us.repo.FindOne(id); err != nil {
    return errors.New("user not found")
  }

  if len(user.Password) < 8 {
    return errors.New("length of password must be at least 8")
  }
  
  return us.repo.Update(id, user)
}