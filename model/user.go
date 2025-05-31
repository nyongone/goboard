package model

import "time"

type ContextKey string
var UserKey = ContextKey("user")

type User struct {
  ID			 			int    			`json:"id"`
  Email		 			string 			`json:"email"`
  Password			string 			`json:"password"`
  Name		 			string 			`json:"name"`
  Birthdate			*time.Time 	`json:"birth_date"`
  RefreshToken 	*string			`json:"refresh_token"`
  CreatedAt 		*time.Time 	`json:"created_at"`
  UpdatedAt 		*time.Time 	`json:"updated_at"`
  DeletedAt 		*time.Time 	`json:"deleted_at"`
}

type UserResponse struct {
  ID						int					`json:"id"`
  Name					string			`json:"name"`
  BirthDate			*time.Time	`json:"birth_date,omitempty"`
  CreatedAt 		*time.Time 	`json:"created_at,omitempty"`
  UpdatedAt 		*time.Time 	`json:"updated_at,omitempty"`
  DeletedAt 		*time.Time 	`json:"deleted_at,omitempty"`
}

type UserRepository interface {
  FindOne(id int) (*User, error)
  FindOneByEmail(email string) (*User, error)
  Create(user *User) error
  Update(id int, user *User) error
  UpdateRefreshToken(id int, token string) error
}

type UserService interface {
  FindOne(id int) (*User, error)
  FindOneByEmail(email string) (*User, error)
  Create(user *User) error
  Update(id int, user *User) error
}