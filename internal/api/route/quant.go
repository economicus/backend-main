package route

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/handler"
)

func SetQuant(router *gin.RouterGroup, handler *handler.QuantHandler) {
	quant := router.Group("/quants")
	{
		quant.GET("", handler.GetAllQuants)
		quant.GET("/quant/:quant_id", handler.GetQuant)

		quant.POST("/quant", handler.CreateQuant)

		quant.PATCH("/quant/:quant_id", handler.UpdateQuant)
		quant.PUT("/quant-option/:quant_id", handler.UpdateQuantOption)

		quant.DELETE("/quant/:quant_id", handler.DeleteQuant)
	}
	lab := router.Group("/lab")
	{
		lab.GET("/list", handler.GetLabList)
		lab.GET("/data/:quant_id", handler.GetLabData)
	}
}
