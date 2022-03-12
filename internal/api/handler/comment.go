package handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/service"
	e "main/internal/core/error"
	"main/internal/core/model"
	"net/http"
)

type CommentHandler struct {
	s *service.CommentService
}

func NewCommentHandler(s *service.CommentService) *CommentHandler {
	return &CommentHandler{
		s: s,
	}
}

// GetCommentsAndReplies returns all comments and replies of a quant model
func (h *CommentHandler) GetCommentsAndReplies(ctx *gin.Context) {
	var q struct {
		QuantID uint `json:"quant_id"`
	}

	if err := ctx.BindQuery(&q); err != nil {
		sendErr(ctx, e.ErrInvalidQuery)
		return
	}

	comments, err := h.s.GetCommentsAndReplies(q.QuantID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// CommentToQuant create a comment
func (h *CommentHandler) CommentToQuant(ctx *gin.Context) {
	var req model.Comment

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	if err = h.s.CreateComment(user.ID, &req); err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// UpdateComment edit a comment
func (h *CommentHandler) UpdateComment(ctx *gin.Context) {
	var req model.Comment

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	if err = h.s.UpdateComment(user.ID, &req); err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// DeleteComment delete a comment
func (h *CommentHandler) DeleteComment(ctx *gin.Context) {
	var req struct {
		CommentID uint `json:"comment_id"`
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

	if err = h.s.DeleteComment(user.ID, req.CommentID); err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
