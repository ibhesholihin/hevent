package models

import "time"

type (
	Event struct {
		ID          uint      `json:"id" gorm:"primaryKey;not null;unique"`
		Event_Name  string    `json:"event_name" gorm:"type:varchar(100);not null;"`
		Description string    `json:"description" gorm:"type:text"`
		ImageURL    string    `json:"image_url" gorm:"type:text;"`
		Start_time  time.Time `json:"start_time" gorm:"not null;"`
		End_time    time.Time `json:"end_time" gorm:"not null;"`
		Location    string    `json:"location" gorm:"type:text"`
		Latitude    string    `json:"latitude" gorm:"type:text"`
		Longitude   string    `json:"longitude" gorm:"type:text"`
		CategoryID  uint      `json:"category_id" gorm:"not null;"`
		Quantity    uint      `json:"quantity" gorm:"not null; default:0"`

		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at" gorm:"default:null"`
		Active    int8      `json:"active" gorm:"type:int;default:1"`

		EventCategory EventCategory `gorm:"foreignKey:CategoryID"`
	}

	EventCategory struct {
		ID        uint      `json:"id" gorm:"primaryKey;not null;unique"`
		Category  string    `json:"kategori" gorm:"type:varchar(50);not null;"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	}

	EventPriceTipe struct {
		ID            uint      `json:"id" gorm:"primaryKey;not null;unique"`
		EventID       uint      `json:"event_id" gorm:"not null"`
		Tipe          string    `json:"tipe" gorm:"type:varchar(50);not null;"`
		Price         uint      `json:"price" gorm:"not null;"`
		Qty           uint      `json:"qty" gorm:"not null; default:0"`
		CreatedAt     time.Time `json:"created_at" gorm:"<-:create"`
		Limit_Time    time.Time `json:"limit_time"`
		IsNormalPrice bool      `json:"is_normal_price" gorm:"default:false"`

		Event Event `gorm:"foreignKey:EventID"`
	}

	EventSponsor struct {
		ID           uint      `json:"id" gorm:"primaryKey;not null;unique"`
		EventID      uint      `json:"event_id" gorm:"not null"`
		Sponsor_Name string    `json:"sponsor" gorm:"type:varchar(100);not null;"`
		CreatedAt    time.Time `json:"created_at" gorm:"<-:create"`

		Event Event `gorm:"foreignKey:EventID"`
	}
)

// Event Request and Response
type (

	//CREATE
	CreateEventReq struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		Quantity      uint   `json:"quantity"`
		EventCategory string `json:"kategori"`
		ImageURL      string `json:"image_url"`

		CategoryID uint   `json:"category_id"`
		Location   string `json:"location" gorm:"type:text"`
		//Tipe Price
		//Price uint `json:"price"`
	}
	CreateEventRes struct {
		ID   uint   `json:"id"`
		Name string `json:"event_name"`
	}

	CreateEventCategoryReq struct {
		Category string `json:"kategori"`
	}

	CreateEventCategoryRes struct {
		ID       uint   `json:"id"`
		Category string `json:"kategori"`
	}

	CreateEventPriceReq struct {
		EventID uint   `json:"event_id"`
		Tipe    string `json:"tipe" `
		Price   uint   `json:"price"`
		Qty     uint   `json:"qty"`
	}

	CreateEventTipeRes struct {
		ID      uint   `json:"id"`
		EventID uint   `json:"event_id"`
		Tipe    string `json:"tipe" `
		Price   uint   `json:"price"`
		Qty     uint   `json:"qty"`
	}

	//GET
	GetEventRes struct {
		ID          uint                `json:"id"`
		Name        string              `json:"name"`
		Description string              `json:"description"`
		Price       uint                `json:"price"`
		Quantity    uint                `json:"quantity"`
		Category    GetEventCategoryRes `json:"category"`
		ImageURL    string              `json:"image_url"`
		CreatedAt   time.Time           `json:"created_at"`
		UpdatedAt   time.Time           `json:"updated_at"`
		DeletedAt   time.Time           `json:"deleted_at"`
	}

	GetEventCategoryRes struct {
		CategoryID   uint   `json:"id"`
		CategoryName string `json:"kategori"`
	}

	GetEventPriceRes struct {
		EventID uint   `json:"event_id"`
		Tipe    string `json:"tipe" `
		Price   uint   `json:"price"`
		Qty     uint   `json:"qty"`
	}

	//UPDATE
	UpdateEventReq struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       uint   `json:"price"`
		Quantity    uint   `json:"quantity"`
		ImageURL    string `json:"image_url"`
		CategoryID  uint   `json:"category_id"`
		Location    string `json:"location" gorm:"type:text"`
	}

	UpdateEventCategoryReq struct {
		ID       uint   `json:"id"`
		Category string `json:"kategori"`
	}

	UpdateEventCategoryRes struct {
		ID       uint   `json:"id"`
		Category string `json:"kategori"`
	}
)

// Event Query Process
type (
	EventQueries struct {
		CategoryID uint
		//Price      []uint
	}
)
