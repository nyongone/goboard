package service

import (
	"errors"
	"go-board-api/internal/util"
	"go-board-api/model"
	"regexp"
)

type signupService struct {
  repo model.UserRepository
}

func NewSignupService(repo model.UserRepository) model.SignupService {
  return &signupService{repo}
}

func (ss *signupService) Signup(user *model.Signup) error {
  if _, err := ss.repo.FindOneByEmail(user.Email); err == nil {
    return errors.New("user already exists")
  }

  var emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
  re := regexp.MustCompile(emailRegex)
  if !re.MatchString(user.Email) {
    return errors.New("email format not matches")
  }

  if len(user.Password) < 8 {
    return errors.New("length of password must be at least 8")
  }

  hashedPassword, err := util.GeneratePassword(user.Password)
  if err != nil {
    return errors.New("error occured while hashing password")
  }

  newUser := &model.User{}
  newUser.Email = user.Email
  newUser.Password = hashedPassword
  newUser.Name = user.Name
  newUser.Birthdate = user.Birthdate

  return ss.repo.Create(newUser)
}