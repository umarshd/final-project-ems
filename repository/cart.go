package repository

import (
	"final-project-ems/entity"

	"gorm.io/gorm"
)

type ICartRepository interface {
	FindAllCart(userId string) ([]entity.Cart, error)
	StoreCart(cart entity.Cart) (entity.Cart, error)
	FindCartById(id int) (entity.Cart, error)
	UpdateCart(cart entity.Cart) (entity.Cart, error)
	StoreCartItem(cartItem entity.CartItem) (entity.CartItem, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db}
}

func (r CartRepository) FindAllCart(userId string) ([]entity.Cart, error) {
	var carts []entity.Cart
	if err := r.db.Debug().Preload("User").Where("user_id = ?", userId).Find(&carts).Error; err != nil {
		return carts, err
	}

	return carts, nil
}

func (r CartRepository) StoreCart(cart entity.Cart) (entity.Cart, error) {
	if err := r.db.Debug().Create(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (r CartRepository) FindCartById(id int) (entity.Cart, error) {
	var cart entity.Cart
	if err := r.db.Debug().Preload("User").Where("id = ?", id).Find(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (r CartRepository) UpdateCart(cart entity.Cart) (entity.Cart, error) {
	if err := r.db.Debug().Where("id = ?", cart.ID).Updates(&cart).Error; err != nil {
		return cart, err
	}

	return cart, nil
}

func (r CartRepository) StoreCartItem(cartItem entity.CartItem) (entity.CartItem, error) {
	if err := r.db.Debug().Create(&cartItem).Error; err != nil {
		return cartItem, err
	}

	return cartItem, nil
}
