package jwt

import (
	"github.com/golang-jwt/jwt"
	e "main/internal/core/error"
	"strconv"
)

func Validate(token, tokenType string) (uint, error) {
	var claims userClaims

	jt, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, e.ErrInvalidTokenAlg
		}
		return []byte(cnf.GetSecret()), nil
	})
	if err != nil {
		return 0, e.ErrInvalidToken
	}
	if !jt.Valid {
		return 0, e.ErrExpiredToken
	}
	if claims.TokenType != tokenType {
		return 0, e.ErrWrongTokenType
	}
	userID, err := strconv.ParseUint(claims.Id, 10, 64)
	if err != nil {
		return 0, e.ErrInvalidUser
	}
	return uint(userID), nil
}
