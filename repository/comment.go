package repository

import (
	"database/sql"
	"go-board-api/model"
)

type commentRepository struct {
  db *sql.DB
}

func NewCommentRepository(db *sql.DB) model.CommentRepository {
  return &commentRepository{db}
}

func (cr *commentRepository) FindOne(id int) (*model.CommentResponse, error) {
  query := `
    SELECT c.id, c.comment, c.created_at, c.updated_at, c.deleted_at, u.id, u.name
    FROM comments AS c
    JOIN users AS u ON u.id = c.author_id
    WHERE c.id = ?
  `
  row := cr.db.QueryRow(query, id)

  comment := &model.CommentResponse{}
  author := &model.CommentAuthor{}
  err := row.Scan(&comment.ID, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt, &author.ID, &author.Name)
  if err != nil {
    return nil, err
  }

  comment.Author = *author

  return comment, nil
}

func (cr *commentRepository) FindAllByPost(post_id int) ([]*model.CommentResponse, error) {
  query := `
    SELECT c.id, c.comment, c.post_id, c.created_at, c.updated_at, c.deleted_at, u.id, u.name
    FROM comments AS c
    JOIN users AS u
    WHERE c.post_id = ?
    ORDER BY c.created_at DESC
  `
  rows, err := cr.db.Query(query, post_id)
  if err != nil {
    return nil, err
  }

  var comments []*model.CommentResponse
  for rows.Next() {
    comment := &model.CommentResponse{}
    author := &model.CommentAuthor{}
    err := rows.Scan(&comment.ID, &comment.Comment, &comment.PostID, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt, &author.ID, &author.Name)
    if err != nil {
      return nil, err
    }

    if comment.DeletedAt != nil { comment.Comment = "삭제된 댓글입니다." }

    comment.Author = *author

    comments = append(comments, comment)
  }

  return comments, nil
}

func (cr *commentRepository) FindAllByAuthor(author_id int) ([]*model.CommentResponse, error) {
  query := `
    SELECT c.id, c.comment, c.post_id, c.created_at, c.updated_at, c.deleted_at, u.id, u.name
    FROM comments AS c
    JOIN users AS u
    WHERE c.author_id = ?
    ORDER BY c.created_at DESC
  `
  rows, err := cr.db.Query(query, author_id)
  if err != nil {
    return nil, err
  }

  var comments []*model.CommentResponse
  for rows.Next() {
    comment := &model.CommentResponse{}
    author := &model.CommentAuthor{}
    err := rows.Scan(&comment.ID, &comment.Comment, &comment.PostID, &comment.CreatedAt, &comment.UpdatedAt, &comment.DeletedAt, &author.ID, &author.Name)
    if err != nil {
      return nil, err
    }

    if comment.DeletedAt != nil { comment.Comment = "삭제된 댓글입니다." }

    comment.Author = *author

    comments = append(comments, comment)
  }

  return comments, nil
}

func (cr *commentRepository) Create(comment *model.Comment) error {
  _, err := cr.db.Exec("INSERT INTO comments (comment, author_id, post_id) VALUES (?, ?, ?)", &comment.Comment, &comment.AuthorID, &comment.PostID)
  if err != nil {
    return err
  }

  return nil
}

func (cr *commentRepository) Delete(id int) error {
  _, err := cr.db.Exec("DELETE FROM comments WHERE id = ?", id)
  if err != nil {
    return err
  }

  return nil
}

func (cr *commentRepository) SoftDelete(id int) error {
  _, err := cr.db.Exec("UPDATE comments SET deleted_at = NOW() WHERE id = ?", id)
  if err != nil {
    return err
  }

  return nil
}