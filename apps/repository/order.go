package repository

import (
	md "github.com/ibhesholihin/hevent/apps/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	OrderRepo interface {
		FindOrCreateCart(userid int64) (md.CartSession, error)
		FindCartItems(sessionid uint) ([]md.CartItem, error)
		AddItemToCart(cartItem md.CartItem) error
	}

	orderRepo struct {
		*gorm.DB
	}
)

func (repo *orderRepo) FindOrCreateCart(userid int64) (md.CartSession, error) {
	cartSession := md.CartSession{}
	if err := repo.DB.FirstOrCreate(&cartSession, md.CartSession{UserID: userid}).Error; err != nil {
		return md.CartSession{}, err
	}
	return cartSession, nil
}

func (repo *orderRepo) FindCartItems(sessionid uint) ([]md.CartItem, error) {
	cartItems := []md.CartItem{}
	if err := repo.DB.Preload("Events").Preload(clause.Associations).Where("session_id = ?", sessionid).Find(&cartItems).Error; err != nil {
		return []md.CartItem{}, err
	}
	return cartItems, nil
}

func (repo *orderRepo) AddItemToCart(cartItem md.CartItem) error {
	result := repo.DB.Debug().Model(md.CartItem{}).Where("session_id = ? AND event_price_id = ?", cartItem.SessionID, cartItem.EventPriceID).Updates(&cartItem)
	if result.RowsAffected != 0 {
		return nil
	}
	if err := repo.DB.Debug().Create(&cartItem).Error; err != nil {
		return err
	}
	return nil
}
