package route

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/handler"
	"main/internal/api/middleware"
)

func SetUser(router *gin.RouterGroup, handler *handler.UserHandler, mid *middleware.AuthMiddleware) {
	router.POST("/register", handler.Register)
	userRoute := router.Group("/users", mid.Authenticate())
	{
		userRoute.GET("", handler.GetAllUsers)
		userRoute.GET("/user", handler.GetUser)
		userRoute.DELETE("/user", handler.DeleteUser)

		userRoute.GET("/profile/:user_id", handler.GetUserProfile)
		userRoute.PATCH("/profile", handler.EditUserProfile)
		userRoute.PUT("/profile-image", handler.UploadUserProfileImage)

		userRoute.GET("/favorite-quants", handler.GetFavoriteQuants)
		userRoute.POST("/favorite-quants", handler.AddToFavoriteQuants)
		userRoute.DELETE("/favorite-quants", handler.DeleteFromFavoriteQuants)

		userRoute.GET("/followings", handler.GetFollowings)
		userRoute.GET("/followers", handler.GetFollowers)
		userRoute.DELETE("/followings", handler.UnfollowUser)
		userRoute.POST("/followings", handler.FollowUser)
	}
}
