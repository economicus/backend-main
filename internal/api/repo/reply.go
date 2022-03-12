package repo

import (
	"gorm.io/gorm"
	"main/internal/core/model"
	"main/internal/pkg/logger"
)

type ReplyRepo struct {
	db *gorm.DB
}

func NewReplyRepository(db *gorm.DB) *ReplyRepo {
	return &ReplyRepo{
		db: db,
	}
}

func (repo *ReplyRepo) GetReply(replyID uint) (*model.Reply, error) {
	var reply model.Reply
	err := repo.db.First(&reply, replyID).Error
	return &reply, err
}

func (repo *ReplyRepo) CreateReply(reply *model.Reply) error {
	if err := repo.db.Create(reply).Error; err != nil {
		logger.Logger.Errorf("error in CreateReply: %v\n", err)
		return err
	}
	return nil
}

func (repo *ReplyRepo) UpdateReply(replyID uint, content string) error {
	if err := repo.db.First(&model.Reply{}, replyID).Update("content", content).Error; err != nil {
		logger.Logger.Errorf("error in UpdateReply: %v\n", err)
		return err
	}
	return nil
}

func (repo *ReplyRepo) DeleteReply(replyID uint) error {
	if err := repo.db.Delete(&model.Reply{}, replyID).Error; err != nil {
		logger.Logger.Errorf("error in DeleteReply: %v\n", err)
		return err
	}
	return nil
}
