package handler

import (
	"github.com/gin-gonic/gin"
	e "main/internal/core/error"
	"main/internal/core/model/table"
	"strconv"
	"strings"
)

func getUserFromContext(ctx *gin.Context) (*table.User, error) {
	setUser, exist := ctx.Get("user")
	if !exist {
		return nil, e.ErrNoUserFound
	}
	user, ok := setUser.(table.User)
	if !ok {
		return nil, e.ErrNoUserFound
	}
	return &user, nil
}

func getFieldsFromContext(ctx *gin.Context) []string {
	field := ctx.Query("fields")
	if field == "" {
		return []string{}
	}
	return strings.Split(field, ",")
}

func extractUserId(ctx *gin.Context) (uint, error) {
	id := ctx.Query("user_id")

	userID64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(userID64), nil
}
