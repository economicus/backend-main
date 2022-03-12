package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main/internal/conf/db/mysql"
	"main/internal/core/model/table"
	"main/internal/pkg/jwt"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	db *mysql.DB
}

func NewAuthMiddleware(db *mysql.DB) *AuthMiddleware {
	return &AuthMiddleware{
		db: db,
	}
}

func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := extractAccessToken(ctx.Request)
		if err != nil {
			abortErr(ctx, err)
			return
		}

		userID, err := jwt.Validate(accessToken, "access")
		if err != nil {
			abortErr(ctx, err)
			return
		}

		user, err := getUserByID(m.db.DB, userID)
		if err != nil {
			abortErr(ctx, err)
			return
		}

		ctx.Set("user", *user)
		ctx.Next()
	}
}

func getUserByID(db *gorm.DB, userID uint) (*table.User, error) {
	var user table.User
	err := db.Preload("Profile").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// extractAccessToken extracts access token from request
func extractAccessToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	data := strings.Split(header, " ")
	if len(data) != 2 {
		return "", fmt.Errorf("error in extractAccessToken while split header")
	}
	return data[1], nil
}
