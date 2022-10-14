package repository

import (
	"final-project-ems/entity"

	"gorm.io/gorm"
)

type IProductRepository interface {
	Store(product entity.Product) (entity.Product, error)
	FindAll() ([]entity.Product, error)
	FindById(id int) (entity.Product, error)
	Find(category int, priceMin string, priceMax string) ([]entity.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r ProductRepository) Store(product entity.Product) (entity.Product, error) {
	if err := r.db.Debug().Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r ProductRepository) FindAll() ([]entity.Product, error) {
	var product []entity.Product
	if err := r.db.Debug().Preload("ProductCategory").Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r ProductRepository) FindById(id int) (entity.Product, error) {
	var product entity.Product
	if err := r.db.Debug().Preload("ProductCategory").Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r ProductRepository) Find(category int, priceMin string, priceMax string) ([]entity.Product, error) {
	var product []entity.Product
	if err := r.db.Debug().Preload("ProductCategory").Where("category_id = ? OR price BETWEEN ? AND ?", category, priceMin, priceMax).Find(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}
