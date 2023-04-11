package service

import (
	"shopper"
	"shopper/pkg/repo"
)

type CommentService struct {
	repo repo.Comment
}

func NewCommentService(repo repo.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) AddComment(itemId, userId int, input shopper.Comment) (int, error) {
	return s.repo.AddComment(itemId, userId, input)
}

func (s *CommentService) GetCommentsByItem(id int) ([]shopper.Comment, error) {
	return s.repo.GetCommentsByItem(id)
}

func (s *CommentService) DeleteComment(userId, id int) error {
	return s.repo.DeleteComment(userId, id)
}

func (s *CommentService) UpdateComment(userId, id int, input shopper.UpdateCommentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.UpdateComment(userId, id, input)
}
