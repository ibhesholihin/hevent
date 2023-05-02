package service

import (
	"context"
	"time"

	md "github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/repository"
)

type (

	//init service contract
	OrderService interface {
		//FindListCategory(c context.Context) ([]md.EventCategory, error)
		//AddCategory(c context.Context, adminReq md.CreateEventCategoryReq) (md.CreateEventCategoryRes, error)
	}

	//init service parameter
	orderService struct {
		stores         *repository.Stores
		contextTimeout time.Duration
	}
)

// Event Category Get List
func (s *eventService) FindListOrder(c context.Context) ([]md.EventCategory, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	listCategory, err := s.stores.Event.FindListCategory()
	if err != nil {
		return []md.EventCategory{}, err
	}
	return listCategory, nil
}

// Create User Cart Order
func (s *eventService) CreateCart(c context.Context, catReq md.CreateEventCategoryReq) (md.CreateEventCategoryRes, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	kategori, err := s.stores.Event.FindOrCreateCategory(catReq.Category)
	if err != nil {
		return md.CreateEventCategoryRes{}, err
	}

	katResp := md.CreateEventCategoryRes{
		ID: kategori.ID, Category: kategori.Category,
	}
	return katResp, nil
}
