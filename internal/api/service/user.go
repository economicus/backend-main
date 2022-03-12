package service

import (
	"fmt"
	"main/internal/api/repo"
	"main/internal/conf/aws"
	"main/internal/core/model"
	"main/internal/core/model/request"
	"main/internal/core/model/table"
	"mime/multipart"
	"time"
)

type UserService struct {
	repo *repo.UserRepo
	aws  *aws.Aws
}

func NewUserService(repo *repo.UserRepo, aws *aws.Aws) *UserService {
	return &UserService{
		repo: repo,
		aws:  aws,
	}
}

func (s *UserService) GetUsers(option *model.Query) (model.Users, error) {
	return s.repo.GetUsers(option)
}

func (s *UserService) GetUser(userID uint, fields []string) (map[string]interface{}, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	user.Profile.ProfileImage = s.aws.GetFilePath(user.Profile.ProfileImage)

	return user.ToMapWithFields(fields), nil
}

func (s *UserService) Register(req *request.RegisterReq) error {
	if err := s.repo.CheckNickname(req.Nickname); err != nil {
		return err
	}
	userID, err := s.repo.CreateUser(req.Email, req.Password)
	if err != nil {
		return err
	}
	return s.repo.CreateProfile(userID, req.Nickname)
}

func (s *UserService) UpdateProfile(userID uint, req *model.Profile) error {
	if req.Nickname != "" {
		err := s.repo.CheckNickname(req.Nickname)
		if err != nil {
			return err
		}
	}

	m := req.ToMap()
	m["updated_at"] = time.Now()

	return s.repo.UpdateProfile(userID, m)
}

func (s *UserService) UploadProfileImage(userID uint, file multipart.File, header *multipart.FileHeader) error {
	filepath := fmt.Sprintf("photos/%s", header.Filename)

	_, err := s.aws.UploadFile(file, header)
	if err != nil {
		return err
	}

	return s.repo.UploadUserProfileImage(userID, filepath)
}

func (s *UserService) DeleteUser(ID uint) error {
	return s.repo.DeleteUser(ID)
}

func (s *UserService) UpdatePassword(userID uint, newPassword string) error {
	return s.repo.UpdatePassword(userID, newPassword)
}

func (s *UserService) GetFollowings(userID uint) (model.Users, error) {
	users, err := s.repo.GetFollowings(userID)
	if err != nil {
		return nil, err
	}

	for idx, user := range users {
		users[idx].Profile.ProfileImage = s.aws.GetFilePath(user.Profile.ProfileImage)
	}

	return users, nil
}

func (s *UserService) GetFollowers(userID uint) (model.Users, error) {
	users, err := s.repo.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	for idx, user := range users {
		users[idx].Profile.ProfileImage = s.aws.GetFilePath(user.Profile.ProfileImage)
	}

	return users, nil
}

func (s *UserService) Follow(userID, followerID uint) error {
	return s.repo.Follow(userID, followerID)
}

func (s *UserService) UnFollow(userID, followingID uint) error {
	return s.repo.UnFollow(userID, followingID)
}

func (s *UserService) GetFavoriteQuants(userID uint) ([]*table.Quant, error) {
	return s.repo.GetFavoriteQuants(userID)
}

func (s *UserService) AddToFavoriteQuants(userID, quantID uint) error {
	return s.repo.AddToFavoriteQuants(userID, quantID)
}

func (s *UserService) DeleteFromFavoriteQuants(userID, quantID uint) error {
	return s.repo.DeleteFromFavoriteQuants(userID, quantID)
}
