package model

import (
	"main/internal/core/model/table"
	"main/internal/pkg/objconv"
)

type Comments []Comment

type Comment struct {
	table.Comment
}

func NewComment(userID, quantID uint, content string) *Comment {
	return &Comment{
		table.Comment{
			UserID:  userID,
			QuantID: quantID,
			Content: content,
		},
	}
}

func (c *Comment) GetID() uint {
	return c.ID
}

func (c *Comment) GetOwnerID() uint {
	return c.UserID
}

func (c *Comment) ToMap() map[string]interface{} {
	return objconv.ToMap(c)
}

func (c *Comment) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(c, fields)
}
