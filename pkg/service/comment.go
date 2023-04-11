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
