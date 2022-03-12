package route

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/handler"
)

func SetAuth(router *gin.RouterGroup, handler *handler.AuthHandler) {
	router.POST("login", handler.LoginInLocal)
	router.DELETE("/logout")
	router.POST("/refresh-token")
}
