package repo

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"main/internal/conf/aws"
	e "main/internal/core/error"
	"main/internal/core/model"
	"main/internal/core/model/table"
	"main/internal/pkg/logger"
	"main/internal/pkg/pwd"
)

type UserRepo struct {
	db  *gorm.DB
	aws *aws.Aws
}

func NewUser(db *gorm.DB, aws *aws.Aws) *UserRepo {
	return &UserRepo{
		db:  db,
		aws: aws,
	}
}

// GetUsers return users
func (repo *UserRepo) GetUsers(option *model.Query) (model.Users, error) {
	var users model.Users

	sql := fmt.Sprintf("select * "+
		"from users as u join profiles as p on u.id = p.user_id "+
		"where u.name like '%%%s%%' or p.nickname like '%%%s%%' "+
		"order by %s limit %d offset %d;",
		option.Keyword, option.Keyword, option.Order, option.PerPage, option.Page*option.PerPage)

	err := repo.db.Preload("Profile").Raw(sql).Find(&users).Error
	if err != nil {
		logger.Logger.Errorf("error in GetUsers: %v\n", err)
		return nil, err
	}

	return users, nil
}

// GetUserByID finds user with user id
func (repo *UserRepo) GetUserByID(userID uint) (*model.User, error) {
	var user model.User

	err := repo.db.Preload("Profile").First(&user, userID).Error
	if err != nil {
		logger.Logger.Errorf("error in GetUserByID: %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepo) CheckNickname(nickname string) error {
	err := repo.db.Where("nickname = ?", nickname).First(&model.Profile{}).Error
	if err == nil {
		return e.ErrDuplicateNickname
	} else if err == gorm.ErrRecordNotFound {
		return nil
	} else {
		logger.Logger.Errorf("error in CheckNickname: %v\n", err)
		return err
	}
}

func (repo *UserRepo) CreateUser(email, password string) (uint, error) {
	user := model.NewUser(email, password, "local")

	if err := repo.db.Create(user).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return 0, e.ErrDuplicateEmail
		} else {
			logger.Logger.Errorf("error in CreateUser: %v\n", err)
			return 0, err
		}
	}

	return user.ID, nil
}

func (repo *UserRepo) CreateProfile(userID uint, nickname string) error {
	profile := model.NewProfile(userID, nickname)

	if err := repo.db.Create(profile).Error; err != nil {
		logger.Logger.Errorf("error in CreateProfile: %v\n", err)
		return err
	}

	return nil
}

// UpdateProfile updates user's profile
func (repo *UserRepo) UpdateProfile(userID uint, data map[string]interface{}) error {
	var profile model.Profile

	err := repo.db.Where("user_id = ?", userID).First(&profile).Updates(data).Error
	if err != nil {
		logger.Logger.Errorf("error in UpdateUserProfile: %v\n", err)
		return err
	}

	return nil
}

// UploadUserProfileImage uploads user's profile image path from s3
func (repo *UserRepo) UploadUserProfileImage(userID uint, filepath string) error {
	var profile model.Profile

	err := repo.db.Model(&profile).Where("user_id = ?", userID).Update("profile_image", filepath).Error
	if err != nil {
		logger.Logger.Errorf("error in UploadUserProfileImage: %v\n", err)
		return err
	}

	return nil
}

// DeleteUser soft-deletes a user
func (repo *UserRepo) DeleteUser(ID uint) error {
	if err := repo.db.First(&model.User{}, ID).Update("user_active", 0).Error; err != nil {
		logger.Logger.Errorf("error in DeleteUser while inactivating user: %v\n", err)
		return err
	}

	if err := repo.db.Delete(&model.User{}, ID).Error; err != nil {
		logger.Logger.Errorf("error in DeleteUser while deleting user: %v\n", err)
		return err
	}

	return nil
}

// UpdatePassword updates a user's password
func (repo *UserRepo) UpdatePassword(userID uint, newPassword string) error {
	newHashedPwd, err := pwd.Hash([]byte(newPassword))

	if err != nil {
		logger.Logger.Errorf("error in UpdatePassword while hashing password: %v\n", err)
		return err
	}

	err = repo.db.First(&model.User{}, userID).Update("password", newHashedPwd).Error
	if err != nil {
		logger.Logger.Errorf("error in UpdatePassword while updating password: %v\n", err)
		return err
	}

	return nil
}

// GetFollowings returns users who are followed by the user
func (repo *UserRepo) GetFollowings(userID uint) (model.Users, error) {
	var followings model.Users

	user, err := repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	err = repo.db.Preload("Profile").Model(&user).Association("Followings").Find(&followings)
	if err != nil {
		logger.Logger.Errorf("error in GetFollowings while following user: %v\n", err)
		return nil, err
	}

	return followings, err
}

// GetFollowers returns users who follow the user
func (repo *UserRepo) GetFollowers(userID uint) (model.Users, error) {
	var users model.Users

	sql := fmt.Sprintf("select * from users where id in (select user_id from followings where following_id = %d)", userID)
	if err := repo.db.Preload("Profile").Raw(sql).Find(&users).Error; err != nil {
		logger.Logger.Errorf("error in GetFollowers: %v\n", err)
		return nil, err
	}

	return users, nil
}

// Follow adds to follower list of the user
func (repo *UserRepo) Follow(userID, followingID uint) error {
	sql := fmt.Sprintf("insert into followings (user_id, following_id) values (%d, %d)", userID, followingID)

	if err := repo.db.Exec(sql).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return e.ErrDuplicateRecord
		} else {
			logger.Logger.Errorf("error in Follow: %v", err)
			return err
		}
	}

	return nil
}

// UnFollow deletes a following from follower list of the user
func (repo *UserRepo) UnFollow(userID, followingID uint) error {
	sql := fmt.Sprintf("delete from followings where user_id = %d and following_id = %d", userID, followingID)

	if err := repo.db.Exec(sql).Error; err != nil {
		logger.Logger.Errorf("error in UnFollow: %v", err)
		return err
	}

	return nil
}

func (repo *UserRepo) GetFavoriteQuants(userID uint) ([]*table.Quant, error) {
	var user model.User

	if err := repo.db.Preload("FavoriteQuants").First(&user, userID).Error; err != nil {
		logger.Logger.Errorf("error in GetFavoriteQuants: %v", err)
		return nil, err
	}

	return user.FavoriteQuants, nil
}

func (repo *UserRepo) AddToFavoriteQuants(userID, quantID uint) error {
	sql := fmt.Sprintf("insert into user_favorite_quants (user_id, quant_id) values (%d, %d)", userID, quantID)

	if err := repo.db.Exec(sql).Error; err != nil {
		logger.Logger.Errorf("error in AddToFavoriteQuants: %v", err)
		return err
	}

	return nil
}

func (repo *UserRepo) DeleteFromFavoriteQuants(userID, quantID uint) error {
	sql := fmt.Sprintf("delete from user_favorite_quants where (user_id, quant_id) = (%d, %d)", userID, quantID)

	if err := repo.db.Exec(sql).Error; err != nil {
		logger.Logger.Errorf("error in DeleteFromFavoriteQuants: %v", err)
		return err
	}

	return nil
}
