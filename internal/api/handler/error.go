package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	e "main/internal/core/error"
	"net/http"
)

type httpError struct {
	Message string `json:"message" example:"Some error message"`
}

func sendErr(ctx *gin.Context, err error) {
	code := e.DetErrCode(err)
	if code == http.StatusInternalServerError {
		sendInternalErr(ctx)
		return
	}
	ctx.JSON(code, httpError{
		Message: err.Error(),
	})
}

func sendErrWithMsg(ctx *gin.Context, err error, msg string) {
	code := e.DetErrCode(err)
	if code == http.StatusInternalServerError {
		sendInternalErr(ctx)
		return
	}
	ctx.JSON(code, httpError{
		Message: fmt.Sprintf("%s: %s", err.Error(), msg),
	})
}

func sendInvalidPathErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, httpError{
		Message: fmt.Sprintf("error while parsing uri: %v", err),
	})
}

func sendJsonParsingErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, httpError{
		Message: fmt.Sprintf("error while parsing json: %v", err),
	})
}

func sendInternalErr(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, httpError{
		Message: "Something went wrong. Please contact economicus members.",
	})
}
