package service

import (
	"errors"
	"go-board-api/model"
)

type postService struct {
  repo model.PostRepository
}

func NewPostService(repo model.PostRepository) model.PostService {
  return &postService{repo}
}

func (ps *postService) FindOne(id int) (*model.PostResponse, error) {
  return ps.repo.FindOne(id)
}

func (ps *postService) FindAll() ([]*model.PostResponse, error) {
  return ps.repo.FindAll()
}

func (ps *postService) FindAllByBoard(board_id int) ([]*model.PostResponse, error) {
  return ps.repo.FindAllByBoard(board_id)
}

func (ps *postService) FindAllByAuthor(author_id int) ([]*model.PostResponse, error) {
  return ps.repo.FindAllByAuthor(author_id)
}

func (ps *postService) Create(post *model.Post) (*int64, error) {
  if len(post.Title) < 5 {
    return nil, errors.New("length of post's title must be at least 5 characters long")
  }

  if len(post.Content) < 5 {
    return nil, errors.New("length of post's content must be at least 5 characters long")
  }

  return ps.repo.Create(post)
}

func (ps *postService) Update(id int, post *model.Post) (*int64, error) {
  if _, err := ps.repo.FindOne(id); err != nil {
    return nil, errors.New("post not found")
  }

  if len(post.Title) < 5 {
    return nil, errors.New("length of post's title must be at least 5 characters long")
  }

  if len(post.Content) < 5 {
    return nil, errors.New("length of post's content must be at least 5 characters long")
  }

  return ps.repo.Update(id, post)
}

func (ps *postService) Delete(id int, soft_delete bool) error {
  if _, err := ps.repo.FindOne(id); err != nil {
    return errors.New("post not found")
  }

  if soft_delete { return ps.repo.SoftDelete(id) }
  return ps.repo.Delete(id)
}