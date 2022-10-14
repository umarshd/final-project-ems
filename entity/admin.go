package entity

import "time"

type Admin struct {
	UID       string    `json:"uid" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null;unique"`
	Email     string    `json:"email" gorm:"type:varchar(100);not null"`
	Password  string    `json:"password" gorm:"type:varchar(60);not null"`
	Nama      string    `json:"nama" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type RequestAdmin struct {
	Email    string `form:"email" binding:"required" validate:"email,required"`
	Password string `form:"password" binding:"required" validate:"required"`
	Nama     string `form:"nama" binding:"required" validate:"required"`
}

type ResponseAdmin struct {
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	Nama      string    `json:"nama"`
	CreatedAt time.Time `json:"created_at"`
}

type RequestLoginAdmin struct {
	Email    string `form:"email" binding:"required" validate:"email,required"`
	Password string `form:"password" binding:"required" validate:"required"`
}

type ResponseAdminLogin struct {
	Token string `json:"token"`
}
