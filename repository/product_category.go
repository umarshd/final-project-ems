package repository

import (
	"final-project-ems/entity"

	"gorm.io/gorm"
)

type IProductCategoryRepository interface {
	Store(productCategory entity.ProductCategory) (entity.ProductCategory, error)
	FindAll() ([]entity.ProductCategory, error)
	// FindById(id int)(entity.ProductCategory, error)
	// Update(productCategory entity.ProductCategory) (entity.ProductCategory, error)
	// Delete(id int) error
}

type ProductCategoryRepository struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{db}
}

func (r ProductCategoryRepository) Store(productCategory entity.ProductCategory) (entity.ProductCategory, error) {
	if err := r.db.Debug().Create(&productCategory).Error; err != nil {
		return productCategory, err
	}

	return productCategory, nil
}

func (r ProductCategoryRepository) FindAll() ([]entity.ProductCategory, error) {
	productCategories := []entity.ProductCategory{}
	if err := r.db.Debug().Find(&productCategories).Error; err != nil {
		return productCategories, err
	}

	return productCategories, nil
}
