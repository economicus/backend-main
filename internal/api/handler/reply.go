package handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/service"
	"main/internal/core/model"
	"net/http"
)

type ReplyHandler struct {
	s *service.ReplyService
}

func NewReplyHandler(s *service.ReplyService) *ReplyHandler {
	return &ReplyHandler{
		s: s,
	}
}

// ReplyToComment create a reply to a comment
func (h *ReplyHandler) ReplyToComment(ctx *gin.Context) {
	var req model.Reply

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	if err = h.s.CreateReply(user.ID, &req); err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// UpdateReply edit a reply
func (h *ReplyHandler) UpdateReply(ctx *gin.Context) {
	var req model.Reply

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	if err = h.s.UpdateReply(user.ID, &req); err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// DeleteReply delete a reply
func (h *ReplyHandler) DeleteReply(ctx *gin.Context) {
	var req struct {
		ReplyID uint `json:"reply_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	if err = h.s.DeleteReply(user.ID, req.ReplyID); err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
