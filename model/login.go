package model

type Login struct {
	Email			string		`json:"email"`
	Password	string		`json:"password"`
}

type LoginService interface {
	FindOneByEmail(email string) (*User, error)
	CreateAccessToken(email string) (string, error)
	CreateRefreshToken() (string, error)
	UpdateRefreshToken(id int, token string) error
}