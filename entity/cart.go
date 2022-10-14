package entity

import "time"

type CartItem struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Cart_id    int       `json:"cart_id" gorm:"type:int;not null"`
	Cart       Cart      `json:"cart" gorm:"foreignKey:Cart_id"`
	Product_id int       `json:"product_id"`
	Product    Product   `json:"product" gorm:"foreignKey:Product_id"`
	Quantity   int       `json:"quantity" gorm:"type:int;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Cart struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	User_id    string    `json:"user_id"`
	User       User      `json:"user" gorm:"foreignKey:User_id"`
	Total      int       `json:"total" gorm:"type:int;not null"`
	IsCheckout bool      `json:"is_checkout" gorm:"type:bool;not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type RequestCart struct {
	User_id string `form:"user_id" binding:"required" validate:"required"`
}

type ResponseCart struct {
	ID         int       `json:"id"`
	User_id    string    `json:"user_id"`
	Total      int       `json:"total"`
	IsCheckout bool      `json:"is_checkout"`
	Created    time.Time `json:"created_at"`
}

type RequestCartItem struct {
	Cart_id    int `form:"cart_id" binding:"required" validate:"required"`
	Product_id int `form:"product_id" binding:"required" validate:"required"`
	Quantity   int `form:"quantity" binding:"required" validate:"required"`
}

type ResponseCartItem struct {
	ID         int       `json:"id"`
	Cart_id    int       `json:"cart_id"`
	Product_id int       `json:"product_id"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
}

type RequestUpdateCart struct {
	Cart_id int `form:"cart_id" binding:"required" validate:"required"`
	Total   int `form:"total" binding:"required" validate:"required"`
}

type RequestUpdateCartStatus struct {
	IsCheckout bool `form:"is_checkout" binding:"required" validate:"required"`
}
