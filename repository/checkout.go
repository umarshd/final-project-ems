package repository

import (
	"final-project-ems/entity"

	"gorm.io/gorm"
)

type ICheckoutRepository interface {
	StoreCheckout(checkout entity.Checkout) (entity.Checkout, error)
}

type CheckoutRepository struct {
	db *gorm.DB
}

func NewCheckoutRepository(db *gorm.DB) *CheckoutRepository {
	return &CheckoutRepository{db}
}

func (r CheckoutRepository) StoreCheckout(checkout entity.Checkout) (entity.Checkout, error) {
	if err := r.db.Debug().Create(&checkout).Error; err != nil {
		return checkout, err
	}

	return checkout, nil
}
