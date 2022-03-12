package service

import (
	"main/internal/api/repo"
	e "main/internal/core/error"
	"main/internal/core/model"
)

type CommentService struct {
	repo *repo.CommentRepo
}

func NewCommentService(repo *repo.CommentRepo) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) GetCommentsAndReplies(quantID uint) (model.Comments, error) {
	return s.repo.GetCommentsAndReplies(quantID)
}

func (s *CommentService) GetComment(commentID uint) (*model.Comment, error) {
	return s.repo.GetComment(commentID)
}

func (s *CommentService) CreateComment(userID uint, comment *model.Comment) error {
	if comment.QuantID == 0 {
		return e.ErrMissingRequest
	}
	newCom := model.NewComment(userID, comment.QuantID, comment.Content)

	return s.repo.CreateComment(newCom)
}

func (s *CommentService) UpdateComment(userID uint, comment *model.Comment) error {
	if comment.QuantID == 0 {
		return e.ErrMissingRequest
	}

	c, err := s.GetComment(comment.ID)
	if err != nil {
		return err
	}

	err = repo.CheckPermission(userID, c)
	if err != nil {
		return err
	}
	return s.repo.UpdateComment(comment.ID, comment.Content)
}

func (s *CommentService) DeleteComment(userID, commentID uint) error {
	c, err := s.GetComment(commentID)
	if err != nil {
		return err
	}

	err = repo.CheckPermission(userID, c)
	if err != nil {
		return err
	}
	return s.repo.DeleteComment(commentID)
}
