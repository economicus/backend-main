package route

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/handler"
)

func SetComment(router *gin.RouterGroup, handler *handler.CommentHandler) {
	comment := router.Group("/comments")
	{
		comment.GET("", handler.GetCommentsAndReplies)
		comment.POST("", handler.CommentToQuant)
		comment.PATCH("", handler.UpdateComment)
		comment.DELETE("", handler.DeleteComment)
	}
}
