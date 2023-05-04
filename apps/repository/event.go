package repository

import (
	md "github.com/ibhesholihin/hevent/apps/models"
	"gorm.io/gorm"
)

type (
	EventRepo interface {
		//Event Category
		FindOrCreateCategory(name string) (md.EventCategory, error)
		FindCategory(id uint) (md.EventCategory, error)
		FindListCategory() ([]md.EventCategory, error)
		UpdateEventCategory(cat md.EventCategory) error

		//Event
		FindListEvent() ([]md.Event, error)
		CreateEvent(event md.Event) (md.Event, error)
		GetEvent(id uint) (md.Event, error)
		UpdateEvent(event md.Event) error

		//Event Price Type
		AddEventPrice(price md.EventPriceTipe) (md.EventPriceTipe, error)
		GetEventPrice(id uint) ([]md.EventPriceTipe, error)
		GetEventPriceByID(id uint) (md.EventPriceTipe, error)
		UpdateEventPrice(price md.EventPriceTipe) (md.EventPriceTipe, error)
	}

	eventRepo struct {
		*gorm.DB
	}
)

// Event Category
func (repo *eventRepo) FindOrCreateCategory(name string) (md.EventCategory, error) {
	category := md.EventCategory{}
	if err := repo.Where(md.EventCategory{Category: name}).FirstOrCreate(&category).Error; err != nil {
		return md.EventCategory{}, err
	}
	return category, nil
}

func (repo *eventRepo) FindCategory(id uint) (md.EventCategory, error) {
	category := md.EventCategory{}
	if err := repo.Where(md.EventCategory{ID: uint(id)}).First(&category).Error; err != nil {
		return md.EventCategory{}, err
	}
	return category, nil
}

func (repo *eventRepo) FindListCategory() ([]md.EventCategory, error) {
	listCategory := []md.EventCategory{}
	if err := repo.Find(&listCategory).Error; err != nil {
		return []md.EventCategory{}, err
	}
	return listCategory, nil
}

func (repo *eventRepo) UpdateEventCategory(cat md.EventCategory) error {
	if err := repo.Save(&cat).Error; err != nil {
		return err
	}
	return nil
}

// Event
func (repo *eventRepo) FindListEvent() ([]md.Event, error) {
	listEvent := []md.Event{}
	if err := repo.Find(&listEvent).Error; err != nil {
		return []md.Event{}, err
	}
	return listEvent, nil
}

func (repo *eventRepo) CreateEvent(event md.Event) (md.Event, error) {
	if err := repo.Create(&event).Error; err != nil {
		return md.Event{}, err
	}
	return event, nil
}

func (repo *eventRepo) GetEvent(id uint) (md.Event, error) {
	event := md.Event{}
	if err := repo.Where("events.id = ?", id).Joins("EventCategory").First(&event).Error; err != nil {
		return md.Event{}, err
	}
	return event, nil
}

func (repo *eventRepo) UpdateEvent(event md.Event) error {
	if err := repo.Save(&event).Error; err != nil {
		return err
	}
	return nil
}

// event price type
func (repo *eventRepo) AddEventPrice(price md.EventPriceTipe) (md.EventPriceTipe, error) {
	if err := repo.Create(&price).Error; err != nil {
		return md.EventPriceTipe{}, err
	}
	return price, nil
}

func (repo *eventRepo) UpdateEventPrice(price md.EventPriceTipe) (md.EventPriceTipe, error) {
	if err := repo.Save(&price).Error; err != nil {
		return md.EventPriceTipe{}, err
	}
	return price, nil
}

func (repo *eventRepo) GetEventPrice(eventid uint) ([]md.EventPriceTipe, error) {
	price := []md.EventPriceTipe{}
	if err := repo.Where("event_price_tipes.event_id = ?", eventid).
		Joins("Events").Find(&price).Error; err != nil {
		return []md.EventPriceTipe{}, err
	}
	return price, nil
}

func (repo *eventRepo) GetEventPriceByID(id uint) (md.EventPriceTipe, error) {
	price := md.EventPriceTipe{}
	if err := repo.Where("event_price_tipes.id = ?", id).
		Joins("Events").First(&price).Error; err != nil {
		return md.EventPriceTipe{}, err
	}
	return price, nil

}
