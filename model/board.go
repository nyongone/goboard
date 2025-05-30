package model

import "time"

type Board struct {
	ID				int					`json:"id"`
	Name			string			`json:"name"`
	Slug			string			`json:"slug"`
	CreatedAt *time.Time	`json:"created_at"`
	UpdatedAt *time.Time	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"deleted_at"`
}

type BoardResponse struct {
	ID				int					`json:"id"`
	Name			string			`json:"name"`
	Slug			string			`json:"slug"`
	CreatedAt *time.Time	`json:"created_at,omitempty"`
	UpdatedAt *time.Time	`json:"updated_at,omitempty"`
	DeletedAt	*time.Time	`json:"deleted_at,omitempty"`
}

type BoardRepository interface {
	FindAll() ([]*Board, error)
}

type BoardService interface {
	FindAll() ([]*Board, error)
}