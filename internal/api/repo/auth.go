package repo

import (
	"gorm.io/gorm"
	e "main/internal/core/error"
	"main/internal/core/model"
	"main/internal/pkg/jwt"
	"main/internal/pkg/logger"
	"main/internal/pkg/pwd"
	"time"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (repo *AuthRepo) RefreshToken(refreshToken string) (string, error) {
	userID, err := jwt.Validate(refreshToken, "refresh")
	if err != nil {
		return "", err
	}
	if err = repo.db.First(&model.User{}, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", e.ErrNoUserFound
		} else if err != nil {
			logger.Logger.Errorf("error in RefreshToken: %v\n", err)
			return "", err
		}
	}
	return jwt.NewAccessToken(userID)
}

// AuthenticateInLocal authenticates user with email and password which stored in local database
func (repo *AuthRepo) AuthenticateInLocal(email, password string) (*jwt.Token, error) {
	var user model.User

	err := repo.db.Where("email = ? AND auth_resource = ?", email, "local").
		First(&user).Update("last_login", time.Now()).Error

	if err == gorm.ErrRecordNotFound {
		return nil, e.ErrNoUserFound
	} else if err != nil {
		logger.Logger.Errorf("error in AuthenticateInLocal: %v\n", err)
		return nil, err
	}

	err = pwd.Compare([]byte(password), user.Password)
	if err != nil {
		return nil, err
	}

	tokens, err := jwt.NewTokenPair(user.ID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
