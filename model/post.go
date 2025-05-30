package model

import "time"

type Post struct {
	ID				int					`json:"id"`
	Title			string			`json:"title"`
	Content		string			`json:"content"`
	AuthorID	int					`json:"author_id"`
	BoardID		int					`json:"board_id"`
	CreatedAt	*time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at"`
	DeletedAt *time.Time	`json:"deleted_at"`
}

type PostResponse struct {
	ID				int									`json:"id"`
	Title			string							`json:"title"`
	Content		string							`json:"content,omitempty"`
	Author		PostAuthor					`json:"author"`
	Board			BoardResponse				`json:"board"`
	CreatedAt *time.Time					`json:"created_at"`
	UpdatedAt	*time.Time					`json:"updated_at,omitempty"`
	DeletedAt *time.Time					`json:"deleted_at,omitempty"`
}

type PostAuthor struct {
	ID				int					`json:"id"`
	Name			string			`json:"name"`
}

type PostRepository interface {
	FindOne(id int) (*PostResponse, error)
	FindAll() ([]*PostResponse, error)
	FindAllByBoard(board_id int) ([]*PostResponse, error)
	FindAllByAuthor(author_id int) ([]*PostResponse, error)
	Create(post *Post) error
	Update(id int, post *Post) error
	Delete(id int) error
	SoftDelete(id int) error
}

type PostService interface {
	FindOne(id int) (*PostResponse, error)
	FindAll() ([]*PostResponse, error)
	FindAllByBoard(board_id int) ([]*PostResponse, error)
	FindAllByAuthor(author_id int) ([]*PostResponse, error)
	Create(post *Post) error
	Update(id int, post *Post) error
	Delete(id int, soft_delete bool) error
}