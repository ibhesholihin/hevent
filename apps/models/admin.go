package models

import "time"

type (
	Admin struct {
		ID            int64  `json:"id" gorm:"type:int;primaryKey;not null;unique"`
		Username      string `json:"username" gorm:"type:varchar(100);not null;unique"`
		Email         string `json:"email" gorm:"type:varchar(100);not null;unique"`
		Password      string `json:"password" gorm:"type:text;not null;"`
		Fullname      string `json:"fullname" gorm:"type:varchar(100);not null"`
		Access_Token  string `json:"access_token" gorm:"type:varchar(200);"`
		Refresh_Token string `json:"refresh_token" gorm:"type:varchar(200);"`

		Active int8 `json:"active" gorm:"type:int;default:1"`

		CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}

	AdminRegReq struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Password string `json:"password"`
	}
	AdminRegRes struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
	}
	AdminLoginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	AdminLoginRes struct {
		AccessToken string `json:"access_token"`
	}

	UpdateAdminReq struct {
		ID        int64     `json:"id"`
		Email     string    `json:"email"`
		Phone     uint      `json:"phone"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
