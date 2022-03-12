package service

import (
	"main/internal/api/repo"
	e "main/internal/core/error"
	"main/internal/core/model"
)

type ReplyService struct {
	repo *repo.ReplyRepo
}

func NewReplyService(repo *repo.ReplyRepo) *ReplyService {
	return &ReplyService{
		repo: repo,
	}
}

func (s *ReplyService) GetReply(replyID uint) (*model.Reply, error) {
	return s.repo.GetReply(replyID)
}

func (s *ReplyService) CreateReply(userID uint, reply *model.Reply) error {
	if reply.CommentID == 0 {
		return e.ErrMissingRequest
	}
	newReply := model.NewReply(userID, reply.CommentID, reply.Content)

	return s.repo.CreateReply(newReply)
}

func (s *ReplyService) UpdateReply(userID uint, req *model.Reply) error {
	reply, err := s.GetReply(req.GetID())
	if err != nil {
		return err
	}

	if err = repo.CheckPermission(userID, reply); err != nil {
		return err
	}
	return s.repo.UpdateReply(reply.GetID(), req.Content)
}

func (s *ReplyService) DeleteReply(userID, replyID uint) error {
	reply, err := s.GetReply(replyID)
	if err != nil {
		return err
	}

	if err = repo.CheckPermission(userID, reply); err != nil {
		return err
	}
	return s.repo.DeleteReply(replyID)
}
