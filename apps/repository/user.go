package repository

import (
	"errors"
	"time"

	"github.com/ibhesholihin/hevent/apps/models"
	"gorm.io/gorm"
)

type (

	//init repo contract
	UserRepo interface {
		CreateUser(user models.User) (models.User, error)
		GetUserList() ([]models.User, error)
		GetUserByUsername(username string) (models.User, error)
		DeleteUserById(id int) error
		UpdateUser(user models.User) (models.User, error)

		FindUserProfileByUID(uid int64) (models.User, error)
		FindUserProfileID(uid int64) (int64, error)
		UpdateProfile(profile models.UserProfile) (models.UserProfile, error)
	}

	userRepo struct {
		*gorm.DB
	}
)

// func Create/Store User
func (repo *userRepo) CreateUser(user models.User) (models.User, error) {
	if err := repo.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// func Get List USer
func (repo *userRepo) GetUserList() ([]models.User, error) {
	var (
		listUser []models.User
		//query      string
		//queryValue []interface{}
	)

	if err := repo.Find(&listUser).Error; err != nil {
		return []models.User{}, err
	}
	return listUser, nil
}

// func Get User By Username
func (repo *userRepo) GetUserByUsername(username string) (models.User, error) {
	user := models.User{}
	if err := repo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// func Update User By ID
func (repo *userRepo) UpdateUser(user models.User) (models.User, error) {
	if err := repo.Model(&user).Updates(user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// func Delete User By ID
func (repo *userRepo) DeleteUserById(id int) error {

	user := models.User{}
	result := repo.Model(&user).Where("id = ?", id).Update("deleted_at", time.Now())

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected. The record probably does not exist")
	}

	return nil
}

// func Get User By ID
func (repo *userRepo) FindUserProfileByUID(uid int64) (models.User, error) {
	user := models.User{}
	if err := repo.Joins("UserProfile").Where("id = ?", uid).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// func Get User By Profile ID
func (repo *userRepo) FindUserProfileID(uid int64) (int64, error) {
	user := models.User{}
	if err := repo.Select("profile_id").Where("uid = ?", uid).First(&user).Error; err != nil {
		return 0, err
	}
	return user.ProfileID, nil
}

// func Update User
func (repo *userRepo) UpdateProfile(profile models.UserProfile) (models.UserProfile, error) {
	if err := repo.Save(&profile).Error; err != nil {
		return models.UserProfile{}, err
	}
	return profile, nil
}
