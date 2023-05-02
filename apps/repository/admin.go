package repository

import (
	"errors"
	"time"

	"github.com/ibhesholihin/hevent/apps/models"
	"gorm.io/gorm"
)

type (
	AdminRepo interface {
		//Get(tx *sql.Tx) ([]models.Book, error)
		CreateAdmin(admin models.Admin) (models.Admin, error)
		GetAdminList() ([]models.Admin, error)
		GetAdminByUsername(username string) (models.Admin, error)
		GetAdminById(id int64) (models.Admin, error)
		UpdateAdmin(admin models.Admin) (models.Admin, error)
		DeleteAdminById(id int) error
	}

	adminRepo struct {
		*gorm.DB
	}
)

// func Create/Store Admin
func (repo *adminRepo) CreateAdmin(admin models.Admin) (models.Admin, error) {
	if err := repo.Create(&admin).Error; err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// func Get ListAdmin
func (repo *adminRepo) GetAdminList() ([]models.Admin, error) {
	var (
		listAdmin []models.Admin
		//query      string
		//queryValue []interface{}
	)

	if err := repo.Find(&listAdmin).Error; err != nil {
		return []models.Admin{}, err
	}
	return listAdmin, nil
}

// func Get Admin By ID
func (repo *adminRepo) GetAdminById(id int64) (models.Admin, error) {
	admin := models.Admin{}
	if err := repo.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// func Get Admin By Username
func (repo *adminRepo) GetAdminByUsername(username string) (models.Admin, error) {
	admin := models.Admin{}
	if err := repo.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// func Update Admin By ID
func (repo *adminRepo) UpdateAdmin(admin models.Admin) (models.Admin, error) {
	if err := repo.Model(&admin).Updates(admin).Error; err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// func Delete Admin By ID
func (repo *adminRepo) DeleteAdminById(id int) error {

	admin := models.Admin{}
	//result := repo.Model(&admin).Where("id = ?", id).Update("deleted_at", time.Now())
	result := repo.Model(&admin).Where("id = ?", id).Updates(models.Admin{Active: 0, DeletedAt: time.Now()})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected. The record probably does not exist")
	}

	return nil
}
