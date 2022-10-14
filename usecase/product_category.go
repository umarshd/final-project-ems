package usecase

import (
	"final-project-ems/entity"
	"final-project-ems/repository"
)

type IProductCategoryUseCase interface {
	CreateProductCategory(productCategory entity.RequestProductCategory) (entity.ProductCategory, error)
	FindAllProductCategory() ([]entity.ProductCategory, error)
}

type ProductCategoryUseCase struct {
	productCategoryRepository repository.IProductCategoryRepository
}

func NewProductCategoryUseCase(productCategoryRepository repository.IProductCategoryRepository) *ProductCategoryUseCase {
	return &ProductCategoryUseCase{productCategoryRepository}
}

func (useCase ProductCategoryUseCase) CreateProductCategory(productCategory entity.RequestProductCategory) (entity.ProductCategory, error) {
	u := entity.ProductCategory{
		Name:        productCategory.Name,
		Description: productCategory.Description,
	}

	productCategoryResponse, err := useCase.productCategoryRepository.Store(u)

	if err != nil {
		return entity.ProductCategory{}, err
	}

	prodCatResponse := entity.ProductCategory{
		ID:          productCategoryResponse.ID,
		Name:        productCategoryResponse.Name,
		Description: productCategoryResponse.Description,
	}

	return prodCatResponse, nil
}

func (useCase ProductCategoryUseCase) FindAllProductCategory() ([]entity.ProductCategory, error) {
	productCategories, err := useCase.productCategoryRepository.FindAll()

	if err != nil {
		return productCategories, err
	}

	return productCategories, nil
}
