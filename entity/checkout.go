package entity

import "time"

type Checkout struct {
	Id        string    `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null;unique"`
	User_id   string    `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:User_id"`
	Cart_id   int       `json:"cart_id"`
	Cart      Cart      `json:"cart" gorm:"foreignKey:Cart_id"`
	Image_url string    `json:"image_url" gorm:"type:text;not null"`
	Total     int       `json:"total" gorm:"type:int;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type RequestCheckout struct {
	User_id   string `form:"user_id" binding:"required" validate:"required"`
	Cart_id   int    `form:"cart_id" binding:"required" validate:"required"`
	Image_url string `form:"image_url" binding:"required" validate:"required"`
}

type ResponseCheckout struct {
	Id string `json:"id"`
}
