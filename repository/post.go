package repository

import (
	"database/sql"
	"go-board-api/model"
)

type postRepository struct {
  db *sql.DB
}

func NewPostRepository(db *sql.DB) model.PostRepository {
  return &postRepository{db}
}

func (pr *postRepository) FindOne(id int) (*model.PostResponse, error) {
  query := `
    SELECT p.id, p.title, p.content, p.created_at, p.updated_at, p.deleted_at, u.id, u.name, b.id, b.name, b.slug
    FROM posts AS p 
    JOIN users AS u ON u.id = p.author_id
    JOIN boards AS b ON b.id = p.board_id
    WHERE p.id = ? AND p.deleted_at IS NULL;
  `
  row := pr.db.QueryRow(query, id)

  post := &model.PostResponse{}
  board := &model.BoardResponse{}
  author := &model.PostAuthor{}
  err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt, &author.ID, &author.Name, &board.ID, &board.Name, &board.Slug)
  if err != nil {
    return nil, err
  }

  post.Author = *author
  post.Board = *board


  return post, nil
}

func (pr *postRepository) FindAll() ([]*model.PostResponse, error) {
  query := `
    SELECT p.id, p.title, p.created_at, p.updated_at, p.deleted_at, u.id, u.name, b.id, b.name, b.slug
    FROM posts AS p 
    JOIN users AS u ON u.id = p.author_id
    JOIN boards AS b ON b.id = p.board_id
    WHERE p.deleted_at IS NULL
    ORDER BY created_at DESC
  `
  rows, err := pr.db.Query(query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var posts = []*model.PostResponse{}
  for rows.Next() {
    post := &model.PostResponse{}
    board := &model.BoardResponse{}
    author := &model.PostAuthor{}
    err := rows.Scan(&post.ID, &post.Title, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt, &author.ID, &author.Name, &board.ID, &board.Name, &board.Slug)
    if err != nil {
      return nil, err
    }

    post.Author = *author
    post.Board = *board

    posts = append(posts, post)
  }

  return posts, err
}

func (pr *postRepository) FindAllByBoard(board_id int) ([]*model.PostResponse, error) {
  query := `
    SELECT p.id, p.title, p.created_at, p.updated_at, p.deleted_at, u.id, u.name, b.id, b.name, b.slug
    FROM posts AS p 
    JOIN users AS u ON u.id = p.author_id
    JOIN boards AS b ON b.id = p.board_id
    WHERE p.board_id = ? AND p.deleted_at IS NULL
    ORDER BY created_at DESC
  `
  rows, err := pr.db.Query(query, board_id)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var posts = []*model.PostResponse{}
  for rows.Next() {
    post := &model.PostResponse{}
    board := &model.BoardResponse{}
    author := &model.PostAuthor{}
    err := rows.Scan(&post.ID, &post.Title, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt, &author.ID, &author.Name, &board.ID, &board.Name, &board.Slug)
    if err != nil {
      return nil, err
    }

    post.Author = *author
    post.Board = *board

    posts = append(posts, post)
  }

  return posts, nil
}

func (pr *postRepository) FindAllByAuthor(author_id int) ([]*model.PostResponse, error) {
  query := `
    SELECT p.id, p.title, p.created_at, p.updated_at, p.deleted_at, u.id, u.name, b.id, b.name, b.slug
    FROM posts AS p 
    JOIN users AS u ON u.id = p.author_id
    JOIN boards AS b ON b.id = p.board_id
    WHERE p.author_id = ? AND p.deleted_at IS NULL
    ORDER BY created_at DESC
  `
  rows, err := pr.db.Query(query, author_id)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var posts = []*model.PostResponse{}
  for rows.Next() {
    post := &model.PostResponse{}
    board := &model.BoardResponse{}
    author := &model.PostAuthor{}
    err := rows.Scan(&post.ID, &post.Title, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt, &author.ID, &author.Name, &board.ID, &board.Name, &board.Slug)
    if err != nil {
      return nil, err 
    }

    post.Author = *author
    post.Board = *board

    posts = append(posts, post)
  }

  return posts, nil
}

func (pr *postRepository) Create(post *model.Post) error {
  _, err := pr.db.Exec("INSERT INTO posts (title, content, author_id, board_id) VALUES (?, ?, ?, ?)", &post.Title, &post.Content, &post.AuthorID, &post.BoardID)
  if err != nil {
    return err
  }

  return nil
}

func (pr *postRepository) Update(id int, post *model.Post) error {
  _, err := pr.db.Exec("UPDATE posts SET title = ?, content = ?, updated_at = NOW() WHERE id = ?", &post.Title, &post.Content, id)
  if err != nil {
    return err
  }

  return nil
}

func (pr *postRepository) Delete(id int) error {
  _, err := pr.db.Exec("DELETE FROM posts WHERE id = ?", id)
  if err != nil {
    return err
  }

  return nil
}

func (pr *postRepository) SoftDelete(id int) error {
  _, err := pr.db.Exec("UPDATE posts SET deleted_at = NOW() WHERE id = ?", id)
  if err != nil {
    return err
  }

  return nil
}