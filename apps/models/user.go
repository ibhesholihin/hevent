package models

import "time"

// Database Table
type (
	User struct {
		ID        int64  `json:"id" gorm:"type:int;primaryKey;not null;unique"`
		Username  string `json:"username" gorm:"type:varchar(20);not null;unique"`
		Password  string `json:"password" gorm:"type:text;not null;"`
		ProfileID int64  `json:"profile_id" gorm:"not null;"`

		Access_Token  string `json:"access_token" gorm:"type:varchar(200);"`
		Refresh_Token string `json:"refresh_token" gorm:"type:varchar(200);"`

		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`

		Active    int8      `json:"active" gorm:"type:int;default:1"`
		DeletedAt time.Time `json:"deleted_at"`

		UserProfile UserProfile `gorm:"foreignKey:ProfileID"`
	}

	UserProfile struct {
		ID        int64     `json:"id" gorm:"type:int;primaryKey;not null;unique"`
		Fullname  string    `json:"fullname" gorm:"type:varchar(100);not null;"`
		Gender    string    `json:"gender" gorm:"type:varchar(20)"`
		Email     string    `json:"email" gorm:"type:varchar(100);not null;unique"`
		Phone     uint      `json:"phone" gorm:"default:0"`
		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserAddress struct {
		ID           uint      `json:"id" gorm:"primaryKey;not null;unique"`
		UserID       int64     `json:"user_id" gorm:"type:int;not null;"`
		AddressLabel string    `json:"address_label" gorm:"type:varchar(30);not null;"`
		AddressLine  string    `json:"address_line" gorm:"type:text;not null;"`
		City         string    `json:"city" gorm:"type:varchar(50);not null;"`
		Province     string    `json:"province" gorm:"type:varchar(50);not null;"`
		PostalCode   uint      `json:"postal_code"`
		Country      string    `json:"country" gorm:"type:varchar(50);not null;"`
		RegionID     string    `json:"region_id" gorm:"type:varchar(10);not null;"`
		IsDefault    bool      `json:"is_default" gorm:"not null;default:false"`
		CreatedAt    time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt    time.Time `json:"updated_at"`
		DeletedAt    time.Time `json:"deleted_at" gorm:"default:null"`
		User         User      `gorm:"foreignKey:UserID"`
	}
)

// User reqeust and Response
// User reqeust and Response
type (
	//User and Profile Models
	UserRegRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Password string `json:"password"`
	}
	UserRegResponse struct {
		UID      int64  `json:"id"`
		Username string `json:"username"`
	}

	UserLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UserLoginResponse struct {
		AccessToken string `json:"access_token"`
	}
	GetUserProfile struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Gender   string `json:"gender"`
		Email    string `json:"email"`
		Phone    uint   `json:"phone"`
	}
	UpdateProfileReq struct {
		ID        int64     `json:"id"`
		Username  string    `json:"username"`
		Fullname  string    `json:"fullname"`
		Gender    string    `json:"gender"`
		Email     string    `json:"email"`
		Phone     uint      `json:"phone"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	//User Address Models
	UserAddressReq struct {
		PostalCode   uint      `json:"postal_code"`
		AddressLabel string    `json:"address_label"`
		AddressLine  string    `json:"address_line"`
		City         string    `json:"city"`
		Province     string    `json:"province"`
		Country      string    `json:"country"`
		RegionID     string    `json:"region_id"`
		IsDefault    bool      `json:"is_default"`
		DeletedAt    time.Time `json:"deleted_at"`
	}
	CreateUserAddressRes struct {
		AddressID uint   `json:"address_id"`
		UserID    string `json:"user_id"`
	}
	UserAddressRes struct {
		ID           uint   `json:"id"`
		AddressLabel string `json:"address_label"`
		AddressLine  string `json:"address_line"`
		City         string `json:"city"`
		Province     string `json:"province"`
		PostalCode   uint   `json:"postal_code"`
		Country      string `json:"country"`
		RegionID     string `json:"region_id"`
		IsDefault    bool   `json:"is_default"`
	}
)
