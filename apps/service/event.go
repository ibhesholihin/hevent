package service

import (
	"context"
	"time"

	md "github.com/ibhesholihin/hevent/apps/models"
	"github.com/ibhesholihin/hevent/apps/repository"
)

type (

	//init service contract
	EventService interface {
		//category event
		FindListCategory(c context.Context) ([]md.EventCategory, error)
		AddCategory(c context.Context, catReq md.CreateEventCategoryReq) (md.CreateEventCategoryRes, error)
		GetEventCategory(c context.Context, catid uint) (md.GetEventCategoryRes, error)

		//event
		FindListEvents(c context.Context) ([]md.Event, error)
		GetEvent(c context.Context, eventid uint) (md.GetEventRes, error)
		CreateEvent(c context.Context, EvReq md.CreateEventReq) (md.CreateEventRes, error)

		//event price tipe
	}

	//init service parameter
	eventService struct {
		stores         *repository.Stores
		contextTimeout time.Duration
	}
)

// Events Get List
func (s *eventService) FindListEvents(c context.Context) ([]md.Event, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	listEvents, err := s.stores.Event.FindListEvent()
	if err != nil {
		return []md.Event{}, err
	}
	return listEvents, nil
}

// Events Get Data
func (s *eventService) GetEvent(c context.Context, eventid uint) (md.GetEventRes, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	event, err := s.stores.Event.GetEvent(eventid)
	if err != nil {
		return md.GetEventRes{}, err
	}
	eventRes := md.GetEventRes{
		ID:          event.ID,
		Name:        event.Event_Name,
		Description: event.Description,
		Quantity:    event.Quantity,
		Category: md.GetEventCategoryRes{
			CategoryID:   event.EventCategory.ID,
			CategoryName: event.EventCategory.Category,
		},
		ImageURL:  event.ImageURL,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
		DeletedAt: event.DeletedAt,
	}
	return eventRes, nil
}

// Event Add
func (s *eventService) CreateEvent(c context.Context, EvReq md.CreateEventReq) (md.CreateEventRes, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	//Store Events
	pd := md.Event{
		Event_Name:  EvReq.Name,
		Description: EvReq.Description,
		Quantity:    EvReq.Quantity,
		ImageURL:    EvReq.ImageURL,
		CategoryID:  EvReq.CategoryID,
	}

	evt, err := s.stores.Event.CreateEvent(pd)
	if err != nil {
		return md.CreateEventRes{}, err
	}
	//Response
	productRes := md.CreateEventRes{
		ID:   evt.ID,
		Name: evt.Event_Name,
	}
	return productRes, nil
}

// Event Category Get List
func (s *eventService) FindListCategory(c context.Context) ([]md.EventCategory, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	listCategory, err := s.stores.Event.FindListCategory()
	if err != nil {
		return []md.EventCategory{}, err
	}
	return listCategory, nil
}

// Event Category Create
func (s *eventService) AddCategory(c context.Context, catReq md.CreateEventCategoryReq) (md.CreateEventCategoryRes, error) {
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

// Events Get Data
func (s *eventService) GetEventCategory(c context.Context, catid uint) (md.GetEventCategoryRes, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	cat, err := s.stores.Event.FindCategory(catid)
	if err != nil {
		return md.GetEventCategoryRes{}, err
	}
	catRes := md.GetEventCategoryRes{
		CategoryID:   cat.ID,
		CategoryName: cat.Category,
	}

	return catRes, nil
}
