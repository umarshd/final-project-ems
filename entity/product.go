package entity

import "time"

type Product struct {
	ID              int             `json:"id" gorm:"primaryKey;autoIncrement"`
	SKU             string          `json:"sku" gorm:"type:varchar(100);not null"`
	Name            string          `json:"name" gorm:"type:varchar(100);not null"`
	Description     string          `json:"description" gorm:"type:text;not null"`
	Price           string          `json:"price" gorm:"type:varchar(20);not null"`
	Quantity        int             `json:"quantity" gorm:"type:int;not null"`
	Image_url       string          `json:"image_url" gorm:"type:text;not null"`
	Category_id     int             `json:"category_id" gorm:"type:int;not null"`
	ProductCategory ProductCategory `json:"product_category" gorm:"foreignKey:Category_id"`
	CreatedAt       time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}

type RequestProduct struct {
	SKU         string `form:"sku" binding:"required" validate:"required"`
	Name        string `form:"name" binding:"required" validate:"required"`
	Description string `form:"description" binding:"required" validate:"required"`
	Price       string `form:"price" binding:"required" validate:"required"`
	Quantity    int    `form:"quantity" binding:"required" validate:"required"`
	Image_url   string `form:"image_url" binding:"required" validate:"required"`
	Category_id int    `form:"category_id" binding:"required" validate:"required"`
}

type ResponseProduct struct {
	ID          int       `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Quantity    int       `json:"quantity"`
	Image_url   string    `json:"image_url"`
	Category_id int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type ResponseGetAllProduct struct {
	Name            string          `json:"name"`
	Image_url       string          `json:"image_url"`
	Price           string          `json:"price"`
	Category_id     int             `json:"category_id"`
	ProductCategory ProductCategory `json:"product_category" gorm:"foreignKey:Category_id"`
}
