package service

import (
	"errors"
	"go-board-api/model"
)

type commentService struct {
  repo model.CommentRepository
}

func NewCommentService(repo model.CommentRepository) model.CommentService {
  return &commentService{repo}
}

func (cs *commentService) FindOne(id int) (*model.CommentResponse, error) {
  return cs.repo.FindOne(id)
}

func (cs *commentService) FindAllByPost(post_id int) ([]*model.CommentResponse, error) {
  return cs.repo.FindAllByPost(post_id)
}

func (cs *commentService) FindAllByAuthor(author_id int) ([]*model.CommentResponse, error) {
  return cs.repo.FindAllByAuthor(author_id)
}

func (cs *commentService) Create(comment *model.Comment) error {
  return cs.repo.Create(comment)
}

func (cs *commentService) Delete(id int, soft_delete bool) error {
  if _, err := cs.repo.FindOne(id); err != nil {
    return errors.New("comment not found")
  }

  if soft_delete { return cs.repo.SoftDelete(id) }
  return cs.repo.Delete(id)
}