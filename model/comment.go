package model

import "time"

type Comment struct {
	ID				int					`json:"id"`
	Comment		string			`json:"comment"`
	AuthorID	int					`json:"author_id"`
	PostID		int					`json:"post_id"`
	CreatedAt	*time.Time	`json:"created_at"`
	UpdatedAt *time.Time	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"deleted_at"`
}

type CommentResponse struct {
	ID				int								`json:"id"`
	Comment		string						`json:"comment"`
	Author		CommentAuthor			`json:"author"`
	PostID		int								`json:"post_id"`
	CreatedAt	*time.Time				`json:"created_at"`
	UpdatedAt *time.Time				`json:"updated_at,omitempty"`
	DeletedAt	*time.Time				`json:"deleted_at,omitempty"`
}

type CommentAuthor struct {
	ID				int					`json:"id"`
	Name			string			`json:"name"`
}

type CommentRepository interface {
	FindOne(id int) (*CommentResponse, error)
	FindAllByPost(post_id int) ([]*CommentResponse, error)
	FindAllByAuthor(author_id int) ([]*CommentResponse, error)
	Create(comment *Comment) error
	Delete(id int) error
	SoftDelete(id int) error
}

type CommentService interface {
	FindOne(id int) (*CommentResponse, error)
	FindAllByPost(post_id int) ([]*CommentResponse, error)
	FindAllByAuthor(author_id int) ([]*CommentResponse, error)
	Create(comment *Comment) error
	Delete(id int, soft_delete bool) error
}