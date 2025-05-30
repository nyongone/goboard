package model

import "time"

type Signup struct {
	Email			string			`json:"email"`
	Password	string			`json:"password"`
	Name			string			`json:"name"`
	Birthdate	*time.Time	`json:"birth_date"`
}

type SignupService interface {
	Signup(user *Signup) error
}