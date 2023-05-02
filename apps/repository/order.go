package repository

import (
	//md "github.com/ibhesholihin/hevent/apps/models"
	"gorm.io/gorm"
)

type (
	OrderRepo interface {
	}

	orderRepo struct {
		*gorm.DB
	}
)
