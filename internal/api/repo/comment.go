package repo

import (
	"gorm.io/gorm"
	"main/internal/core/model"
	"main/internal/pkg/logger"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (repo *CommentRepo) GetCommentsAndReplies(quantID uint) (model.Comments, error) {
	var comments model.Comments

	err := repo.db.Preload("Replies").Where("quant_id = ?", quantID).Find(&comments).Error
	if err != nil {
		logger.Logger.Errorf("error in GetCommentsAndReplies: %v\n", err)
		return nil, err
	}
	return comments, nil
}

func (repo *CommentRepo) GetComment(commentID uint) (*model.Comment, error) {
	var comment model.Comment

	if err := repo.db.First(&comment, commentID).Error; err != nil {
		logger.Logger.Errorf("error in GetComment: %v\n", err)
		return nil, err
	}
	return &comment, nil
}

func (repo *CommentRepo) CreateComment(comment *model.Comment) error {
	if err := repo.db.Create(comment).Error; err != nil {
		logger.Logger.Errorf("error in CreateComment: %v\n", err)
		return err
	}
	return nil
}

func (repo *CommentRepo) UpdateComment(commentID uint, content string) error {
	if err := repo.db.First(&model.Comment{}, commentID).Update("content", content).Error; err != nil {
		logger.Logger.Errorf("error in UpdateComment: %v\n", err)
		return err
	}
	return nil
}

func (repo *CommentRepo) DeleteComment(commentID uint) error {
	if err := repo.db.Delete(&model.Comment{}, commentID).Error; err != nil {
		logger.Logger.Errorf("error in DeleteComment: %v\n", err)
		return err
	}
	return nil
}
