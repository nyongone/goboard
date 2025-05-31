package repository

import (
	"database/sql"
	"go-board-api/model"
)

type userRepository struct {
  db *sql.DB
}

func NewUserRepository(db *sql.DB) model.UserRepository {
  return &userRepository{db}
}

func (ur *userRepository) FindOne(id int) (*model.User, error) {
  row := ur.db.QueryRow("SELECT name, birth_date, created_at, updated_at, deleted_at FROM users WHERE id = ?", id)
  
  user := &model.User{}
  err := row.Scan(&user.Name, &user.Birthdate, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
  if err != nil {
    return nil, err
  }

  return user, nil
}

func (ur *userRepository) FindOneByEmail(email string) (*model.User, error) {
  row := ur.db.QueryRow("SELECT id, email, password, name, birth_date, created_at, updated_at, deleted_at FROM users WHERE email = ?", email)

  user := &model.User{}
  err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Birthdate, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
  if err != nil {
    return nil, err
  }

  return user, nil
}

func (ur *userRepository) Create(user *model.User) error {
  _, err := ur.db.Exec("INSERT INTO users (email, password, name, birth_date) VALUES (?, ?, ?, ?)", &user.Email, &user.Password, &user.Name, &user.Birthdate)
  if err != nil {
    return err
  }

  return nil
}

func (ur *userRepository) Update(id int, user *model.User) error {
  _, err := ur.db.Exec("UPDATE users SET name = ?, password = ?, updated_at = NOW() WHERE id = ?", &user.Name, &user.Password, id)
  if err != nil {
    return err
  }

  return nil
}

func (ur *userRepository) UpdateRefreshToken(id int, token string) error {
  _, err := ur.db.Exec("UPDATE users SET refresh_token = ? WHERE id = ?", token, id)
  if err != nil {
    return err
  }

  return nil
}