package handler

import (
	"final-project-ems/entity"
	"final-project-ems/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUseCase usecase.IProductUseCase
}

func NewProductHandler(productUseCase usecase.IProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase}
}

// Create Product godoc
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer Token"
// @Param body body entity.RequestProduct true "Create Product"
// @Success 201 {object} entity.ResponseProduct
// @Failure 400 {object} entity.ErrorResponse
// @Router /products [post]
func (handler ProductHandler) CreateProduct(c echo.Context) error {
	var req entity.RequestProduct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	product, err := handler.ProductUseCase.CreateProduct(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Product created successfully",
		Data: entity.ResponseProduct{
			ID:          product.ID,
			SKU:         product.SKU,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			Image_url:   product.Image_url,
			Category_id: product.Category_id,
			CreatedAt:   time.Now(),
		},
	})
}

// Get All Product godoc
// @Summary Get All Product
// @Description Get All Product
// @Tags Product
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.ResponseProduct
// @Failure 400 {object} entity.ErrorResponse
// @Router /products [get]
func (handler ProductHandler) FindAllProduct(c echo.Context) error {
	product, err := handler.ProductUseCase.FindAllProduct()

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	if len(product) < 1 {
		return c.JSON(http.StatusNotFound, entity.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Product not found",
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Product found successfully",
		Data:    product,
	})
}

// Get Product By ID godoc
// @Summary Get Product By ID
// @Description Get Product By ID
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} entity.ResponseProduct
// @Failure 400 {object} entity.ErrorResponse
// @Router /products/{id} [get]
func (handler ProductHandler) FindByIdProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := handler.ProductUseCase.FindByIdProduct(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Product found successfully",
		Data:    product,
	})
}

// Filter Product godoc
// @Summary Filter Product
// @Description Filter Product
// @Tags Product
// @Accept  json
// @Produce  json
// @Param category query int false "Category ID"
// @Param priceMin query int false "Price Min"
// @Param priceMax query int false "Price Max"
// @Success 200 {object} entity.ResponseProduct
// @Failure 400 {object} entity.ErrorResponse
// @Router /products/filter [get]
func (handler ProductHandler) FindProduct(c echo.Context) error {
	category, _ := strconv.Atoi(c.FormValue("category"))
	priceMin := c.FormValue("priceMin")
	priceMax := c.FormValue("priceMax")

	product, err := handler.ProductUseCase.FindProduct(category, priceMin, priceMax)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Product found successfully",
		Data:    product,
	})
}
