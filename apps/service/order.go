package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	md "github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/repository"
	"github.com/ibhesholihin/hevent/utils/paygate"
)

type (

	//init service contract
	OrderService interface {
		FindOrCreateCart(c context.Context, userid int64) (interface{}, error)
		AddItemToCart(cartItemReq md.AddItemToCartReq) error

		FindListOrder(c context.Context) ([]md.EventCategory, error)
		TestPayment(c context.Context) (string, string, error)
	}

	//init service parameter
	orderService struct {
		stores         *repository.Stores
		contextTimeout time.Duration
		payment        paygate.PayService
	}
)

// Get List order
func (s *orderService) FindListOrder(c context.Context) ([]md.EventCategory, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	listCategory, err := s.stores.Event.FindListCategory()
	if err != nil {
		return []md.EventCategory{}, err
	}
	return listCategory, nil
}

// Transaction Cart
func (s *orderService) FindOrCreateCart(c context.Context, userid int64) (interface{}, error) {
	cartSession, err1 := s.stores.Order.FindOrCreateCart(userid)
	if err1 != nil {
		return md.CartSession{}, err1
	}
	cartItems, err2 := s.stores.Order.FindCartItems(cartSession.ID)
	if err2 != nil {
		return md.CartSession{}, err2
	}
	resultCartItems := []md.CartItemRes{}

	for _, item := range cartItems {
		cartSession.Total = cartSession.Total + (item.Quantity * item.EventPrice.Price)
		dataAppend := md.CartItemRes{
			ID:        item.ID,
			SessionID: item.SessionID,
			EventID:   item.EventPrice.EventID,
			Quantity:  item.Quantity,
			UpdatedAt: item.UpdatedAt,
			Event:     item.EventPrice.Event,
		}
		resultCartItems = append(resultCartItems, dataAppend)
	}

	cartResult := md.CartSessionRes{
		ID:        cartSession.ID,
		UserUID:   cartSession.UserID,
		Total:     cartSession.Total,
		UpdatedAt: cartSession.UpdatedAt,
		CartItems: resultCartItems,
	}
	return cartResult, nil
}

func (s *orderService) AddItemToCart(cartItemReq md.AddItemToCartReq) error {
	cartItem := md.CartItem{
		SessionID:    cartItemReq.SessionID,
		EventPriceID: cartItemReq.EventPriceID,
		Quantity:     cartItemReq.Quantitty,
	}
	err := s.stores.Order.AddItemToCart(cartItem)
	if err != nil {
		return err
	}
	return nil
}

// Test Payment
func (s *orderService) TestPayment(c context.Context) (string, string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	payID, payUrl, err := s.payment.GeneratePayReq(ctx, "john doe", "admin@demo.id",
		"114432", 2000000)

	fmt.Println("URl :", payUrl)

	if err != nil {
		return "failed", "", errors.New("error")
	}
	return payID, payUrl, nil
}
