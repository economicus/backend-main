package model

import (
	"main/internal/core/model/table"
	"main/internal/pkg/objconv"
)

type Replies []Reply

type Reply struct {
	table.Reply
}

func NewReply(userID, commentID uint, content string) *Reply {
	return &Reply{
		table.Reply{
			UserID:    userID,
			CommentID: commentID,
			Content:   content,
		},
	}
}

func (r *Reply) GetID() uint {
	return r.ID
}

func (r *Reply) GetOwnerID() uint {
	return r.UserID
}

func (r *Reply) ToMap() map[string]interface{} {
	return objconv.ToMap(r)
}

func (r *Reply) ToMapWithFields(fields []string) map[string]interface{} {
	return objconv.ToMapWithFields(r, fields)
}
