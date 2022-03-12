package jwt

import (
	"github.com/golang-jwt/jwt"
	"main/internal/pkg/logger"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewTokenPair(id uint) (*Token, error) {
	at, err := NewAccessToken(id)
	if err != nil {
		return nil, err
	}
	rt, err := NewRefreshToken(id)
	if err != nil {
		return nil, err
	}
	return &Token{
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}

func NewAccessToken(id uint) (string, error) {
	return createToken(id, "access")
}

func NewRefreshToken(id uint) (string, error) {
	return createToken(id, "refresh")
}

func createToken(id uint, t string) (string, error) {
	var claims *userClaims

	claims = NewUserClaims(id, t)
	token := jwt.NewWithClaims(cnf.Alg, claims)
	signedToken, err := token.SignedString([]byte(cnf.GetSecret()))

	if err != nil {
		logger.Logger.Errorf("error while creating jwt token: %v", err)
		return "", err
	}

	return signedToken, nil
}
