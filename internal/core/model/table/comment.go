package table

import "gorm.io/gorm"

type Comment struct {
	gorm.Model `json:"-" structs:"-"`
	UserID     uint    `json:"user_id"`
	QuantID    uint    `json:"quant_id"`
	Content    string  `gorm:"type:text;column:content" json:"content"`
	Replies    []Reply `gorm:"constraint:OnDelete:CASCADE;" json:"replies"`
}
