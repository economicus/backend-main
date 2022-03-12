package service

import (
	"main/internal/api/repo"
	"main/internal/pkg/jwt"
)

type AuthService struct {
	repo *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) LoginInLocal(email, password string) (*jwt.Token, error) {
	return s.repo.AuthenticateInLocal(email, password)
}

func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	return s.repo.RefreshToken(refreshToken)
}
