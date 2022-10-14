package usecase

import (
	"final-project-ems/entity"
	"final-project-ems/repository"
)

type IProductUseCase interface {
	CreateProduct(product entity.RequestProduct) (entity.Product, error)
	FindAllProduct() ([]entity.Product, error)
	FindByIdProduct(id int) (entity.Product, error)
	FindProduct(category int, priceMin string, priceMax string) ([]entity.Product, error)
}

type ProductUseCase struct {
	productRepository repository.IProductRepository
}

func NewProductUseCase(productRepository repository.IProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepository}
}

func (useCase ProductUseCase) CreateProduct(product entity.RequestProduct) (entity.Product, error) {
	u := entity.Product{
		SKU:         product.SKU,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Image_url:   product.Image_url,
		Category_id: product.Category_id,
	}

	productResponse, err := useCase.productRepository.Store(u)

	if err != nil {
		return entity.Product{}, err
	}

	prodResponse := entity.Product{
		ID:          productResponse.ID,
		SKU:         productResponse.SKU,
		Name:        productResponse.Name,
		Description: productResponse.Description,
		Price:       productResponse.Price,
		Image_url:   productResponse.Image_url,
		Quantity:    productResponse.Quantity,
		Category_id: productResponse.Category_id,
	}

	return prodResponse, nil
}

func (useCase ProductUseCase) FindAllProduct() ([]entity.Product, error) {
	product, err := useCase.productRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (useCase ProductUseCase) FindByIdProduct(id int) (entity.Product, error) {
	product, err := useCase.productRepository.FindById(id)

	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (useCase ProductUseCase) FindProduct(category int, priceMin string, priceMax string) ([]entity.Product, error) {
	product, err := useCase.productRepository.Find(category, priceMin, priceMax)

	if err != nil {
		return []entity.Product{}, err
	}

	return product, nil
}
