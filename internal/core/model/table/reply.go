package table

import "gorm.io/gorm"

type Reply struct {
	gorm.Model `json:"-"`
	CommentID  uint   `json:"comment_id"`
	UserID     uint   `json:"user_id"`
	Content    string `gorm:"type:text;column:content" json:"content"`
}
