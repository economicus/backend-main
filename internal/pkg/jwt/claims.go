package jwt

import (
	"github.com/golang-jwt/jwt"
	err "main/internal/core/error"
	"strconv"
	"time"
)

type userClaims struct {
	jwt.StandardClaims
	TokenType string
}

func NewUserClaims(id uint, t string) *userClaims {
	var claims userClaims
	var d time.Duration

	if t == "access" {
		d = cnf.AccessDuration
	} else {
		d = cnf.RefreshDuration
	}

	claims.ExpiresAt = time.Now().Add(d).Unix()
	claims.Id = strconv.Itoa(int(id))
	claims.TokenType = t
	return &claims
}

func (u *userClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return err.ErrInvalidToken
	}
	if u.Id == "0" {
		return err.ErrInvalidUser
	}
	return nil
}
