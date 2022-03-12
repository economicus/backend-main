package route

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/handler"
)

func SetReply(router *gin.RouterGroup, handler *handler.ReplyHandler) {
	reply := router.Group("/replies")
	{
		reply.POST("", handler.ReplyToComment)
		reply.PATCH("", handler.UpdateReply)
		reply.DELETE("", handler.DeleteReply)
	}
}
