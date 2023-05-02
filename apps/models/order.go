package models

import "time"

type (
	CartSession struct {
		ID        uint      `json:"id" gorm:"primaryKey;not null;unique"`
		UserID    string    `json:"user_id" gorm:"type:int;not null;"`
		Total     uint      `json:"total" gorm:"not null;default:0"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		User      User      `gorm:"foreignKey:UserID"`
	}

	CartItem struct {
		ID           uint           `json:"id" gorm:"primaryKey;not null;unique"`
		SessionID    uint           `json:"session_id" gorm:"not_null;"`
		EventPriceID uint           `json:"event_price_id" gorm:"not null;unique"`
		Quantity     uint           `json:"quantity"`
		CreatedAt    time.Time      `json:"created_at" gorm:"<-:create"`
		UpdatedAt    time.Time      `json:"updated_at"`
		CartSession  CartSession    `gorm:"foreignKey:SessionID"`
		EventPrice   EventPriceTipe `gorm:"foreignKey:EventPriceID"`
	}

	Order struct {
		ID             uint           `json:"id" gorm:"primaryKey;not null;unique"`
		UserID         int64          `json:"user_id" gorm:"type:int;not null;"`
		PaymentID      uint           `json:"payment_id" gorm:"not null;"`
		Total          uint           `json:"total" gorm:"not null;default:0"`
		Status         string         `json:"status" gorm:"not null;"`
		CreatedAt      time.Time      `json:"created_at" gorm:"<-:create"`
		UpdatedAt      time.Time      `json:"updated_at"`
		User           User           `gorm:"foreignKey:UserID"`
		PaymentDetails PaymentDetails `gorm:"foreignKey:PaymentID"`
	}

	OrderItem struct {
		ID           uint           `json:"id" gorm:"primaryKey;not null;unique"`
		OrderID      uint           `json:"order_id" gorm:"not_null;"`
		EventPriceID uint           `json:"event_price_id" gorm:"not null;unique"`
		Quantity     uint           `json:"quantity"`
		CreatedAt    time.Time      `json:"created_at" gorm:"<-:create"`
		UpdatedAt    time.Time      `json:"updated_at"`
		Order        Order          `gorm:"foreignKey:OrderID"`
		EventPrice   EventPriceTipe `gorm:"foreignKey:EventPriceID"`
	}

	PaymentDetails struct {
		ID         uint      `json:"id" gorm:"primaryKey;not null;unique"`
		Ammount    uint      `json:"ammount" gorm:"not null;default:0"`
		ReceiptURL string    `json:"receipt_url" gorm:"type:text"`
		CreatedAt  time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt  time.Time `json:"updated_at"`

		PaymentMethod string `json:"payment_method" gorm:"type:text;not null;"`
	}
)

type (
	CartSessionRes struct {
		ID        uint          `json:"id"`
		UserUID   int64         `json:"user_uid"`
		Total     uint          `json:"total"`
		UpdatedAt time.Time     `json:"updated_at"`
		CartItems []CartItemRes `json:"cart_items"`
	}
	CartItemRes struct {
		ID        uint      `json:"id"`
		SessionID uint      `json:"session_id"`
		EventID   uint      `json:"event_id"`
		Quantity  uint      `json:"quantity"`
		UpdatedAt time.Time `json:"updated_at"`
		Event     Event     `gorm:"foreignKey:EventID"`
	}

	AddItemToCartReq struct {
		SessionID uint `json:"session_id"`
		ProductID uint `json:"product_id"`
		Quantitty uint `json:"quantity"`
	}

	CreateOrderReq struct {
		UserID        uint `json:"user_id"`
		UserPaymentID uint `json:"user_payment_id"`
	}

	CreateOrderItemsReq struct {
		ID        uint `json:"id"`
		ProductID uint `json:"product_id"`
		Quantitty uint `json:"quantity"`
	}

	GetOrdersRes struct {
		ID        uint      `json:"id"`
		UserID    int64     `json:"user_id"`
		AddressID uint      `json:"address_id"`
		PaymentID uint      `json:"payment_id"`
		Total     uint      `json:"total"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	GetOrderRes struct {
		ID             uint                   `json:"id"`
		UserUID        string                 `json:"user_uid"`
		AddressID      uint                   `json:"address_id"`
		PaymentID      uint                   `json:"payment_id"`
		Total          uint                   `json:"total"`
		Status         string                 `json:"status"`
		CreatedAt      time.Time              `json:"created_at"`
		UpdatedAt      time.Time              `json:"updated_at"`
		User           GetOrderUserRes        `json:"user_details"`
		UserAddress    GetOrderUserAddressRes `json:"shipment_address"`
		PaymentDetails GetOrderUserPaymentRes `json:"payment_details"`
		OrderItems     []GetOrderItemRes
	}
	GetOrderItemRes struct {
		ID        uint      `json:"id"`
		OrderID   uint      `json:"order_id"`
		EventID   uint      `json:"event_id"`
		Quantity  uint      `json:"quantity"`
		UpdatedAt time.Time `json:"updated_at"`
		Event     Event     `json:"events"`
	}
	GetOrderUserRes struct {
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Phone    uint   `json:"phone"`
	}
	GetOrderUserAddressRes struct {
		ID           uint   `json:"id"`
		AddressLabel string `json:"address_label"`
		AddressLine  string `json:"address_line"`
		City         string `json:"city"`
		Province     string `json:"province"`
		PostalCode   uint   `json:"postal_code"`
		Country      string `json:"country"`
		RegionID     string `json:"region_id"`
	}
	GetOrderUserPaymentRes struct {
		ID            uint      `json:"id"`
		UserPaymentID uint      `json:"user_payment_id"`
		Ammount       uint      `json:"ammount"`
		ReceiptURL    string    `json:"receipt_url"`
		PaymentType   string    `json:"payment_type"`
		Provider      string    `json:"provider"`
		AccountNumber uint      `json:"account_number"`
		Exp           time.Time `json:"exp"`
	}
	ReceiptURLReq struct {
		ReceiptURL string `json:"receipt_url"`
	}
)
