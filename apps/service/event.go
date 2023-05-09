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
		UpdateEventCategory(c context.Context, catReq md.EventCategory, catid uint) (md.EventCategory, error)

		//event
		FindListEvents(c context.Context) ([]md.Event, error)
		GetEvent(c context.Context, eventid uint) (md.GetEventRes, error)
		CreateEvent(c context.Context, EvReq md.CreateEventReq) (md.CreateEventRes, error)
		UpdateEvent(c context.Context, EvReq md.UpdateEventReq, eventid uint) error

		//event price tipe
		FindListPrice(c context.Context, eventid uint) ([]md.EventPriceTipe, error)
		AddEventPrice(c context.Context, EvReq md.EventPriceTipe) (md.EventPriceTipe, error)
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

	dateFormat := "2006-01-02 15:04:05"
	startEvent, _ := time.Parse(dateFormat, EvReq.Start_time)
	endTime, _ := time.Parse(dateFormat, EvReq.End_time)

	//Store Events
	evt := md.Event{
		Event_Name:  EvReq.Event_Name,
		Description: EvReq.Description,
		Quantity:    EvReq.Quantity,
		ImageURL:    EvReq.ImageURL,
		CategoryID:  EvReq.CategoryID,

		Location:  EvReq.Location,
		Latitude:  EvReq.Latitude,
		Longitude: EvReq.Longitude,

		Start_time: startEvent,
		End_time:   endTime,
	}

	evt, err := s.stores.Event.CreateEvent(evt)
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

// Event Update
func (s *eventService) UpdateEvent(c context.Context, EvReq md.UpdateEventReq, eventid uint) error {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	event := md.Event{
		ID:          eventid,
		CategoryID:  EvReq.CategoryID,
		Event_Name:  EvReq.Name,
		Description: EvReq.Description,
		Quantity:    EvReq.Quantity,
		Location:    EvReq.Location,
	}

	err := s.stores.Event.UpdateEvent(event)
	if err != nil {
		return err
	}
	return nil
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

// Event Category Update
func (s *eventService) UpdateEventCategory(c context.Context, catReq md.EventCategory, catid uint) (md.EventCategory, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	cat := md.EventCategory{
		ID:       catid,
		Category: catReq.Category,
	}

	err := s.stores.Event.UpdateEventCategory(cat)
	if err != nil {
		return md.EventCategory{}, err
	}
	return cat, nil

}

// Event Category Get
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

// Event Price Get
func (s *eventService) FindListPrice(c context.Context, eventid uint) ([]md.EventPriceTipe, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	price, err := s.stores.Event.GetEventPrice(eventid)
	if err != nil {
		return []md.EventPriceTipe{}, err
	}

	return price, nil

}

// Event Price Add
func (s *eventService) AddEventPrice(c context.Context, req md.EventPriceTipe) (md.EventPriceTipe, error) {
	_, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()

	price, err := s.stores.Event.AddEventPrice(req)
	if err != nil {
		return md.EventPriceTipe{}, err
	}

	Resp := md.EventPriceTipe{
		ID:      price.ID,
		EventID: price.EventID,
		Tipe:    price.Tipe,
		Price:   price.Price,
	}
	return Resp, nil
}
