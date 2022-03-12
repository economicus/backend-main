package middleware

import (
	"github.com/gin-gonic/gin"
	e "main/internal/core/error"
	"net/http"
)

func abortErr(ctx *gin.Context, err error) {
	code := e.DetErrCode(err)

	if code == http.StatusInternalServerError {
		abortInternalErr(ctx)
		return
	}
	ctx.JSON(code, gin.H{
		"message": err.Error(),
	})
	ctx.Abort()
}

func abortInternalErr(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong. Please contact economicus members.",
	})
	ctx.Abort()
}
