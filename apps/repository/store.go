package repository

import "gorm.io/gorm"

type Stores struct {
	DB    *gorm.DB
	Admin AdminRepo
	User  UserRepo
	Event EventRepo
	Order OrderRepo
}

func NewRepository(db *gorm.DB) *Stores {
	return &Stores{
		DB:    db,
		Admin: &adminRepo{db},
		User:  &userRepo{db},
		Event: &eventRepo{db},
		Order: &orderRepo{db},
	}
}
