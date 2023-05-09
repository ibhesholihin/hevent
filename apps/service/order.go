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
		FindListOrder(c context.Context) ([]md.EventCategory, error)
		//AddCategory(c context.Context, adminReq md.CreateEventCategoryReq) (md.CreateEventCategoryRes, error)

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

// Create User Cart Order
func (s *orderService) CreateCart(c context.Context, catReq md.CreateEventCategoryReq) (md.CreateEventCategoryRes, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	kategori, err := s.stores.Event.FindOrCreateCategory(catReq.Category)
	if err != nil {
		return md.CreateEventCategoryRes{}, err
	}

	cart := md.CreateEventCategoryRes{
		ID: kategori.ID, Category: kategori.Category,
	}
	return cart, nil
}

// Test Payment
func (s *orderService) TestPayment(c context.Context) (string, string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	payID, payUrl, err := s.payment.GeneratePayReq(ctx, "ibhe", "redaksi@herald.id", "114444", 2000000)

	fmt.Println("URl :", payUrl)

	if err != nil {
		return "failed", "", errors.New("error")
	}
	return payID, payUrl, nil
}
